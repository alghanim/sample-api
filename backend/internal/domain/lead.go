package domain

import "time"

// Lead captures inbound interest for Thunder experiences.
type Lead struct {
	ID        string    `json:"id"`
	FullName  string    `json:"fullName"`
	Email     string    `json:"email"`
	Company   string    `json:"company"`
	EventType string    `json:"eventType"`
	Budget    string    `json:"budget"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}

// LeadInput represents the payload used to create a lead.
type LeadInput struct {
	FullName  string `json:"fullName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Company   string `json:"company"`
	EventType string `json:"eventType"`
	Budget    string `json:"budget"`
	Message   string `json:"message"`
}
