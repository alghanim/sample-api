package main

import (
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/thunder-org/thunder-events/internal/config"
	"github.com/thunder-org/thunder-events/internal/http/handler"
	"github.com/thunder-org/thunder-events/internal/http/middleware"
	"github.com/thunder-org/thunder-events/internal/http/router"
	"github.com/thunder-org/thunder-events/internal/repository/pocketbase"
	"github.com/thunder-org/thunder-events/internal/service"
)

func main() {
	cfg := config.MustLoad()

	e := echo.New()

	repo := pocketbase.NewRepository(cfg.PocketBase)
	experienceService := service.NewExperienceService(repo)
	eventHandler := handler.NewEventHandler(experienceService)
	keycloakMiddleware := middleware.NewKeycloakJWT(cfg.Keycloak)

	router.Register(e, eventHandler, keycloakMiddleware, cfg.AllowedOrigins)

	address := cfg.ServerPort
	if !strings.HasPrefix(address, ":") {
		address = ":" + address
	}

	e.Logger.Fatal(e.Start(address))
}
