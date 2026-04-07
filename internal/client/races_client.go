package client

import (
	"fmt"
	"net/url"
	"time"
)

// RacesClient calls races-api endpoints.
type RacesClient struct {
	*Base
}

// NewRacesClient creates a RacesClient.
func NewRacesClient(baseURL, internalKey string, timeout time.Duration) *RacesClient {
	return &RacesClient{Base: newBase(baseURL, internalKey, timeout)}
}

// NewRacesClientNoAuth returns a RacesClient without the internal key.
func NewRacesClientNoAuth(baseURL string, timeout time.Duration) *RacesClient {
	return &RacesClient{Base: newBase(baseURL, "", timeout)}
}

// ── Response models ────────────────────────────────────────────────────────

type Race struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	RaceDate  string  `json:"race_date"`
	EndDate   *string `json:"end_date"`
	Location  string  `json:"location"`
	Region    string  `json:"region"`
	Country   string  `json:"country"`
	Distances string  `json:"distances"`
	URL       string  `json:"url"`
	Source    string  `json:"source"`
	SyncedAt  string  `json:"synced_at"`
}

type RaceSyncResult struct {
	Source  string `json:"source"`
	New     int    `json:"new"`
	Skipped int    `json:"skipped"`
	Error   string `json:"error"`
}

// ── Endpoints ──────────────────────────────────────────────────────────────

// GetHealth calls GET /health.
func (c *RacesClient) GetHealth() (*Response, error) {
	return c.Get("/health")
}

// ListRaces calls GET /api/v1/races with optional query params.
func (c *RacesClient) ListRaces(region string, upcoming bool, limit int) (*Response, error) {
	q := url.Values{}
	if region != "" {
		q.Set("region", region)
	}
	if !upcoming {
		q.Set("upcoming", "false")
	}
	if limit > 0 {
		q.Set("limit", fmt.Sprintf("%d", limit))
	}
	path := "/api/v1/races"
	if len(q) > 0 {
		path += "?" + q.Encode()
	}
	return c.Get(path)
}

// TriggerSync calls POST /api/v1/races/sync.
func (c *RacesClient) TriggerSync() (*Response, error) {
	return c.Post("/api/v1/races/sync", nil)
}
