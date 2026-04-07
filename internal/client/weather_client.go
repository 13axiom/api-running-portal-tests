package client

import (
	"fmt"
	"net/url"
	"time"
)

// WeatherClient calls weather-api endpoints.
type WeatherClient struct {
	*Base
}

// NewWeatherClient creates a WeatherClient.
func NewWeatherClient(baseURL, internalKey string, timeout time.Duration) *WeatherClient {
	return &WeatherClient{Base: newBase(baseURL, internalKey, timeout)}
}

// ── Response models ────────────────────────────────────────────────────────

type City struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type CurrentWeather struct {
	Temperature float64 `json:"temperature"`
	WindSpeed   float64 `json:"wind_speed"`
	WeatherCode int     `json:"weathercode"`
	SyncedAt    string  `json:"synced_at"`
}

type WeatherResponse struct {
	City    City           `json:"city"`
	Current CurrentWeather `json:"current"`
}

type SyncResult struct {
	City    string `json:"city"`
	New     int    `json:"new"`
	Updated int    `json:"updated"`
	Error   string `json:"error"`
}

type AirQualitySnapshot struct {
	CityName string  `json:"city_name"`
	AQI      int     `json:"aqi"`
	PM25     float64 `json:"pm2_5"`
	PM10     float64 `json:"pm10"`
	SyncedAt string  `json:"synced_at"`
}

// ── Endpoints ──────────────────────────────────────────────────────────────

// GetHealth calls GET /health.
func (c *WeatherClient) GetHealth() (*Response, error) {
	return c.Get("/health")
}

// GetCities calls GET /api/v1/cities.
func (c *WeatherClient) GetCities() (*Response, error) {
	return c.Get("/api/v1/cities")
}

// GetWeather calls GET /api/v1/weather/{city}.
func (c *WeatherClient) GetWeather(city string) (*Response, error) {
	return c.Get(fmt.Sprintf("/api/v1/weather/%s", url.PathEscape(city)))
}

// TriggerSync calls POST /api/v1/sync.
func (c *WeatherClient) TriggerSync() (*Response, error) {
	return c.Post("/api/v1/sync", nil)
}

// GetAllAirQuality calls GET /api/v1/air.
func (c *WeatherClient) GetAllAirQuality() (*Response, error) {
	return c.Get("/api/v1/air")
}

// GetCityAirQuality calls GET /api/v1/air/{city}.
func (c *WeatherClient) GetCityAirQuality(city string) (*Response, error) {
	return c.Get(fmt.Sprintf("/api/v1/air/%s", url.PathEscape(city)))
}

// TriggerAirSync calls POST /api/v1/air/sync.
func (c *WeatherClient) TriggerAirSync() (*Response, error) {
	return c.Post("/api/v1/air/sync", nil)
}

// ── Unauthenticated variants (for auth tests) ──────────────────────────────

// GetWeatherNoAuth returns a WeatherClient without the internal key.
func NewWeatherClientNoAuth(baseURL string, timeout time.Duration) *WeatherClient {
	return &WeatherClient{Base: newBase(baseURL, "", timeout)}
}
