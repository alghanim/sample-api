package service

import (
	"context"

	"github.com/thunder-org/thunder-events/internal/domain"
	"github.com/thunder-org/thunder-events/internal/repository/pocketbase"
)

// ExperienceService orchestrates business logic for sporting events.
type ExperienceService interface {
	ListEvents(ctx context.Context) ([]domain.Event, error)
	GetEvent(ctx context.Context, id string) (*domain.Event, error)
	CreateEvent(ctx context.Context, input domain.EventInput) (*domain.Event, error)
	CreateLead(ctx context.Context, input domain.LeadInput) (*domain.Lead, error)
}

type experienceService struct {
	repo pocketbase.Repository
}

// NewExperienceService instantiates the default ExperienceService.
func NewExperienceService(repo pocketbase.Repository) ExperienceService {
	return &experienceService{repo: repo}
}

func (s *experienceService) ListEvents(ctx context.Context) ([]domain.Event, error) {
	return s.repo.ListEvents(ctx)
}

func (s *experienceService) GetEvent(ctx context.Context, id string) (*domain.Event, error) {
	return s.repo.GetEvent(ctx, id)
}

func (s *experienceService) CreateEvent(ctx context.Context, input domain.EventInput) (*domain.Event, error) {
	return s.repo.CreateEvent(ctx, input)
}

func (s *experienceService) CreateLead(ctx context.Context, input domain.LeadInput) (*domain.Lead, error) {
	return s.repo.CreateLead(ctx, input)
}
