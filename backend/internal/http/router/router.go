package router

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/thunder-org/thunder-events/internal/http/handler"
	thunderMW "github.com/thunder-org/thunder-events/internal/http/middleware"
)

// Register wires middleware and routes.
func Register(e *echo.Echo, eventHandler *handler.EventHandler, keycloak *thunderMW.KeycloakJWT, allowedOrigins []string) {
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())

	corsCfg := echoMiddleware.CORSConfig{
		AllowOrigins: allowedOrigins,
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		MaxAge:       int((12 * time.Hour).Seconds()),
	}
	e.Use(echoMiddleware.CORSWithConfig(corsCfg))

	api := e.Group("/api")
	api.GET("/events", eventHandler.ListEvents)
	api.GET("/events/:id", eventHandler.GetEvent)
	api.POST("/leads", eventHandler.CreateLead)

	secured := api.Group("")
	secured.Use(keycloak.Handler)
	secured.POST("/events", eventHandler.CreateEvent)
}
