package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/thunder-org/thunder-events/internal/domain"
	"github.com/thunder-org/thunder-events/internal/service"
)

// EventHandler exposes Thunder experience endpoints.
type EventHandler struct {
	service service.ExperienceService
}

// NewEventHandler builds a handler backed by ExperienceService.
func NewEventHandler(service service.ExperienceService) *EventHandler {
	return &EventHandler{service: service}
}

// ListEvents returns every public event.
func (h *EventHandler) ListEvents(c echo.Context) error {
	events, err := h.service.ListEvents(c.Request().Context())
	if err != nil {
		return respondError(c, http.StatusBadGateway, err)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": events})
}

// GetEvent returns a single event by id.
func (h *EventHandler) GetEvent(c echo.Context) error {
	id := c.Param("id")
	event, err := h.service.GetEvent(c.Request().Context(), id)
	if err != nil {
		return respondError(c, http.StatusBadGateway, err)
	}

	return c.JSON(http.StatusOK, event)
}

// CreateEvent allows authenticated admins to create new events.
func (h *EventHandler) CreateEvent(c echo.Context) error {
	var payload domain.EventInput
	if err := c.Bind(&payload); err != nil {
		return respondError(c, http.StatusBadRequest, err)
	}

	event, err := h.service.CreateEvent(c.Request().Context(), payload)
	if err != nil {
		return respondError(c, http.StatusBadGateway, err)
	}

	return c.JSON(http.StatusCreated, event)
}

// CreateLead tracks marketing submissions from the site.
func (h *EventHandler) CreateLead(c echo.Context) error {
	var payload domain.LeadInput
	if err := c.Bind(&payload); err != nil {
		return respondError(c, http.StatusBadRequest, err)
	}

	lead, err := h.service.CreateLead(c.Request().Context(), payload)
	if err != nil {
		return respondError(c, http.StatusBadGateway, err)
	}

	return c.JSON(http.StatusCreated, lead)
}

func respondError(c echo.Context, status int, err error) error {
	return c.JSON(status, echo.Map{"error": err.Error()})
}
