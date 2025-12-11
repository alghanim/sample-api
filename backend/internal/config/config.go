package config

import (
	"fmt"

	"github.com/caarlos0/env/v9"
)

// Config holds runtime configuration for the Thunder backend.
type Config struct {
	ServerPort     string           `env:"SERVER_PORT" envDefault:"8080"`
	PublicBaseURL  string           `env:"PUBLIC_BASE_URL" envDefault:"http://localhost:8080"`
	AllowedOrigins []string         `env:"CORS_ALLOWED_ORIGINS" envSeparator:"," envDefault:"*"`
	Keycloak       KeycloakConfig   `envPrefix:"KEYCLOAK_"`
	PocketBase     PocketBaseConfig `envPrefix:"POCKETBASE_"`
}

// KeycloakConfig keeps Keycloak integration settings.
type KeycloakConfig struct {
	BaseURL      string `env:"BASE_URL,required"`
	Realm        string `env:"REALM,required"`
	ClientID     string `env:"CLIENT_ID,required"`
	ClientSecret string `env:"CLIENT_SECRET,required"`
}

// PocketBaseConfig stores PocketBase connection details.
type PocketBaseConfig struct {
	BaseURL        string `env:"BASE_URL,required"`
	AdminToken     string `env:"ADMIN_TOKEN,required"`
	PublicFilesURL string `env:"FILES_BASE_URL" envDefault:""`
}

// Load reads configuration from environment variables.
func Load() (Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return Config{}, fmt.Errorf("parse env: %w", err)
	}

	if cfg.PocketBase.PublicFilesURL == "" {
		cfg.PocketBase.PublicFilesURL = cfg.PocketBase.BaseURL
	}

	return cfg, nil
}

// MustLoad wraps Load and panics if parsing fails.
func MustLoad() Config {
	cfg, err := Load()
	if err != nil {
		panic(err)
	}
	return cfg
}
