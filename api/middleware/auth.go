package middleware

import (
	"net/http"

	"github.com/Nerzal/gocloak/v8"
	"github.com/labstack/echo"
)

func KeycloakMiddleware(client gocloak.GoCloak, realm string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get the token from the Authorization header
			authHeader := c.Request().Header.Get("Authorization")

			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing authorization header")
			}

			// We assume the header is in the format `Bearer {token}`
			token := authHeader[len("Bearer "):]

			// Verify the token with Keycloak
			result, err := client.RetrospectToken(c.Request().Context(), token, "client-id", "client-secret", realm)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			// If the token is not active, reject the request
			if !*result.Active {
				return echo.NewHTTPError(http.StatusUnauthorized, "token is not active")
			}

			// If the token is valid, call the next handler
			return next(c)
		}
	}
}
