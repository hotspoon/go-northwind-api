package config

import (
	"os"

	"github.com/rs/zerolog/log"
)

type AppConfig struct {
	JWTSecret   string
	Port        string
	DBPath      string
	Environment string // diganti dari "Env"
	APIVersion  string // diganti dari "APIVer"
}

// LoadConfig membaca env vars dan memberi default
func LoadConfig() *AppConfig {
	cfg := &AppConfig{
		JWTSecret:   os.Getenv("JWT_SECRET"),
		Port:        os.Getenv("PORT"),
		DBPath:      os.Getenv("DB_PATH"),
		Environment: os.Getenv("GO_ENV"),
		APIVersion:  os.Getenv("API_VERSION"),
	}

	// Validasi & default
	if cfg.JWTSecret == "" {
		log.Fatal().Msg("JWT_SECRET is required")
	}
	if cfg.Port == "" {
		cfg.Port = "8080"
	}
	if cfg.DBPath == "" {
		cfg.DBPath = "northwind.db"
	}
	if cfg.Environment == "" {
		cfg.Environment = "development"
	}
	if cfg.APIVersion == "" {
		cfg.APIVersion = "v1"
	}

	return cfg
}

// ==== Implementasi interface routes.ConfigView ====

func (c *AppConfig) Env() string {
	return c.Environment
}

func (c *AppConfig) APIVer() string {
	return c.APIVersion
}
