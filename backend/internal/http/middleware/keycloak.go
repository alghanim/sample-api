package middleware

import (
	"net/http"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/labstack/echo/v4"

	"github.com/thunder-org/thunder-events/internal/config"
)

// KeycloakJWT validates access tokens issued by Keycloak.
type KeycloakJWT struct {
	client       *gocloak.GoCloak
	realm        string
	clientID     string
	clientSecret string
}

// NewKeycloakJWT wires the middleware with the provided configuration.
func NewKeycloakJWT(cfg config.KeycloakConfig) *KeycloakJWT {
	return &KeycloakJWT{
		client:       gocloak.NewClient(cfg.BaseURL),
		realm:        cfg.Realm,
		clientID:     cfg.ClientID,
		clientSecret: cfg.ClientSecret,
	}
}

// Handler returns the actual Echo middleware.
func (m *KeycloakJWT) Handler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing Authorization header")
		}

		if !strings.HasPrefix(strings.ToLower(authHeader), "bearer ") {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid Authorization header")
		}

		token := strings.TrimSpace(authHeader[7:])
		result, err := m.client.RetrospectToken(c.Request().Context(), token, m.clientID, m.clientSecret, m.realm)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		if result == nil || result.Active == nil || !*result.Active {
			return echo.NewHTTPError(http.StatusUnauthorized, "token expired or inactive")
		}

		return next(c)
	}
}
