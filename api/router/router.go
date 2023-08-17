package router

import (
	"alghanim/mediacmsAPI/api/handler"
	"alghanim/mediacmsAPI/api/middleware"
	"alghanim/mediacmsAPI/repository"
	"alghanim/mediacmsAPI/service"
	"database/sql"

	"github.com/Nerzal/gocloak/v8"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo, db *sql.DB) {
	// Initialize Keycloak client
	client := gocloak.NewClient("https://your-keycloak-server.com")

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)

	// Setup routes with Keycloak middleware
	e.GET("/users/:id", userHandler.Get, middleware.KeycloakMiddleware(client, "your-realm"))

}
