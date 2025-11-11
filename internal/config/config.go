package config

import (
	"os"
	"strconv"
)

// Config holds application configuration.
type Config struct {
	Server ServerConfig
	SWAPI  SWAPIConfig
}

// ServerConfig holds server-related configuration.
type ServerConfig struct {
	Port string
}

// SWAPIConfig holds SWAPI-related configuration.
type SWAPIConfig struct {
	BaseURL  string
	PageSize int // Number of items per page to return to clients
}

// Load loads configuration from environment variables with defaults.
func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", ":6969"),
		},
		SWAPI: SWAPIConfig{
			BaseURL:  getEnv("SWAPI_BASE_URL", "https://swapi.dev/api"),
			PageSize: getEnvAsInt("SWAPI_PAGE_SIZE", 15),
		},
	}
}

// getEnv gets environment variable or returns default value.
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt gets environment variable as int or returns default value.
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}
