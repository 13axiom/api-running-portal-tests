// Package config loads test environment configuration.
package config

import (
	"os"
	"time"
)

// Config holds URLs and credentials for both backend APIs.
type Config struct {
	// WeatherAPI base URL (default: http://localhost:8080)
	WeatherAPIURL string

	// RacesAPI base URL (default: http://localhost:8081)
	RacesAPIURL string

	// Shared internal API key (X-Internal-Key header)
	InternalAPIKey string

	// How long to wait for a single HTTP request
	RequestTimeout time.Duration
}

// Load reads config from environment variables.
// All variables have sensible local-dev defaults.
func Load() *Config {
	return &Config{
		WeatherAPIURL:  getEnv("WEATHER_API_URL", "http://localhost:8080"),
		RacesAPIURL:    getEnv("RACES_API_URL",   "http://localhost:8081"),
		InternalAPIKey: getEnv("INTERNAL_API_KEY", ""),
		RequestTimeout: 10 * time.Second,
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
