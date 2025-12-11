package pocketbase

import (
	"context"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/thunder-org/thunder-events/internal/config"
	"github.com/thunder-org/thunder-events/internal/domain"
)

// Repository exposes read/write helpers backed by PocketBase.
type Repository interface {
	ListEvents(ctx context.Context) ([]domain.Event, error)
	GetEvent(ctx context.Context, id string) (*domain.Event, error)
	CreateEvent(ctx context.Context, input domain.EventInput) (*domain.Event, error)
	CreateLead(ctx context.Context, input domain.LeadInput) (*domain.Lead, error)
}

type repository struct {
	client     *resty.Client
	filesURL   string
	eventsPath string
	leadsPath  string
}

// NewRepository wires a PocketBase-backed repository.
func NewRepository(cfg config.PocketBaseConfig) Repository {
	client := resty.New().
		SetBaseURL(cfg.BaseURL).
		SetHeader("Content-Type", "application/json").
		SetAuthToken(cfg.AdminToken)

	return &repository{
		client:     client,
		filesURL:   cfg.PublicFilesURL,
		eventsPath: "/api/collections/events/records",
		leadsPath:  "/api/collections/leads/records",
	}
}

func (r *repository) ListEvents(ctx context.Context) ([]domain.Event, error) {
	var result listResponse[eventRecord]
	if _, err := r.client.R().
		SetContext(ctx).
		SetResult(&result).
		SetQueryParam("sort", "-startDate").
		Get(r.eventsPath); err != nil {
		return nil, fmt.Errorf("list events: %w", err)
	}
	return mapEvents(result.Items, r.filesURL), nil
}

func (r *repository) GetEvent(ctx context.Context, id string) (*domain.Event, error) {
	record := eventRecord{}
	if _, err := r.client.R().
		SetContext(ctx).
		SetResult(&record).
		Get(fmt.Sprintf("%s/%s", r.eventsPath, id)); err != nil {
		return nil, fmt.Errorf("get event: %w", err)
	}

	event := mapEvent(record, r.filesURL)
	return &event, nil
}

func (r *repository) CreateEvent(ctx context.Context, input domain.EventInput) (*domain.Event, error) {
	payload := map[string]any{
		"title":        input.Title,
		"subtitle":     input.Subtitle,
		"description":  input.Description,
		"location":     input.Location,
		"venue":        input.Venue,
		"status":       input.Status,
		"startDate":    input.StartDate,
		"endDate":      input.EndDate,
		"primaryImage": input.PrimaryImage,
		"gallery":      input.Gallery,
		"tags":         input.Tags,
	}

	record := eventRecord{}
	if _, err := r.client.R().
		SetContext(ctx).
		SetBody(payload).
		SetResult(&record).
		Post(r.eventsPath); err != nil {
		return nil, fmt.Errorf("create event: %w", err)
	}

	event := mapEvent(record, r.filesURL)
	return &event, nil
}

func (r *repository) CreateLead(ctx context.Context, input domain.LeadInput) (*domain.Lead, error) {
	payload := map[string]any{
		"fullName":  input.FullName,
		"email":     input.Email,
		"company":   input.Company,
		"eventType": input.EventType,
		"budget":    input.Budget,
		"message":   input.Message,
	}

	record := leadRecord{}
	if _, err := r.client.R().
		SetContext(ctx).
		SetBody(payload).
		SetResult(&record).
		Post(r.leadsPath); err != nil {
		return nil, fmt.Errorf("create lead: %w", err)
	}

	lead := mapLead(record)
	return &lead, nil
}

type listResponse[T any] struct {
	Page       int `json:"page"`
	PerPage    int `json:"perPage"`
	TotalItems int `json:"totalItems"`
	Items      []T `json:"items"`
}

type eventRecord struct {
	ID             string    `json:"id"`
	CollectionID   string    `json:"collectionId"`
	CollectionName string    `json:"collectionName"`
	Title          string    `json:"title"`
	Subtitle       string    `json:"subtitle"`
	Description    string    `json:"description"`
	Location       string    `json:"location"`
	Venue          string    `json:"venue"`
	Status         string    `json:"status"`
	Tags           []string  `json:"tags"`
	StartDate      time.Time `json:"startDate"`
	EndDate        time.Time `json:"endDate"`
	Primary        string    `json:"primaryImage"`
	Gallery        []string  `json:"gallery"`
	Slug           string    `json:"slug"`
}

type leadRecord struct {
	ID        string    `json:"id"`
	FullName  string    `json:"fullName"`
	Email     string    `json:"email"`
	Company   string    `json:"company"`
	EventType string    `json:"eventType"`
	Budget    string    `json:"budget"`
	Message   string    `json:"message"`
	Created   time.Time `json:"created"`
}

func mapEvents(records []eventRecord, baseURL string) []domain.Event {
	items := make([]domain.Event, 0, len(records))
	for _, rec := range records {
		items = append(items, mapEvent(rec, baseURL))
	}
	return items
}

func mapEvent(rec eventRecord, baseURL string) domain.Event {
	event := domain.Event{
		ID:            rec.ID,
		Slug:          rec.Slug,
		Title:         rec.Title,
		Subtitle:      rec.Subtitle,
		Description:   rec.Description,
		Location:      rec.Location,
		Venue:         rec.Venue,
		StartDate:     rec.StartDate,
		EndDate:       rec.EndDate,
		Status:        rec.Status,
		Tags:          rec.Tags,
		GalleryImages: make([]string, 0, len(rec.Gallery)),
	}

	filesCollection := rec.CollectionName
	if filesCollection == "" {
		filesCollection = rec.CollectionID
	}

	if rec.Primary != "" {
		event.PrimaryImage = fmt.Sprintf("%s/api/files/%s/%s/%s", baseURL, filesCollection, rec.ID, rec.Primary)
	}

	for _, g := range rec.Gallery {
		event.GalleryImages = append(event.GalleryImages, fmt.Sprintf("%s/api/files/%s/%s/%s", baseURL, filesCollection, rec.ID, g))
	}
	return event
}

func mapLead(rec leadRecord) domain.Lead {
	return domain.Lead{
		ID:        rec.ID,
		FullName:  rec.FullName,
		Email:     rec.Email,
		Company:   rec.Company,
		EventType: rec.EventType,
		Budget:    rec.Budget,
		Message:   rec.Message,
		CreatedAt: rec.Created,
	}
}
