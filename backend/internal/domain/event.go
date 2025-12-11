package domain

import "time"

// Event models a Thunder sporting experience.
type Event struct {
	ID            string    `json:"id"`
	Slug          string    `json:"slug"`
	Title         string    `json:"title"`
	Subtitle      string    `json:"subtitle"`
	Description   string    `json:"description"`
	Location      string    `json:"location"`
	Venue         string    `json:"venue"`
	StartDate     time.Time `json:"startDate"`
	EndDate       time.Time `json:"endDate"`
	Status        string    `json:"status"`
	PrimaryImage  string    `json:"primaryImage"`
	GalleryImages []string  `json:"galleryImages"`
	Tags          []string  `json:"tags"`
}

// EventInput collects attributes required to create/update an event.
type EventInput struct {
	Title        string    `json:"title" validate:"required"`
	Subtitle     string    `json:"subtitle"`
	Description  string    `json:"description"`
	Location     string    `json:"location" validate:"required"`
	Venue        string    `json:"venue"`
	StartDate    time.Time `json:"startDate" validate:"required"`
	EndDate      time.Time `json:"endDate" validate:"required"`
	Status       string    `json:"status"`
	PrimaryImage string    `json:"primaryImage"`
	Gallery      []string  `json:"gallery"`
	Tags         []string  `json:"tags"`
}
