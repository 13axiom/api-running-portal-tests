// Package auth_test verifies that protected endpoints reject unauthenticated requests.
//
// This is a security regression suite — if any of these tests fail it means
// a protected route is accidentally exposed without auth.
package auth_test

import (
	"net/http"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/13axiom/api-running-portal-tests/internal/client"
	"github.com/13axiom/api-running-portal-tests/internal/config"
	"github.com/joho/godotenv"
)

type AuthSuite struct {
	suite.Suite
	cfg         *config.Config
	weatherNoAuth *client.WeatherClient
	racesNoAuth   *client.RacesClient
}

func (s *AuthSuite) BeforeAll() {
	_ = godotenv.Load("../../.env")
	s.cfg = config.Load()
	// Clients intentionally created WITHOUT the internal key
	s.weatherNoAuth = client.NewWeatherClientNoAuth(s.cfg.WeatherAPIURL, s.cfg.RequestTimeout)
	s.racesNoAuth   = client.NewRacesClientNoAuth(s.cfg.RacesAPIURL, s.cfg.RequestTimeout)
}

func TestAuth(t *testing.T) {
	suite.RunSuite(t, new(AuthSuite))
}

// ── Weather API ────────────────────────────────────────────────────────────

func (s *AuthSuite) TestWeatherCitiesRequiresAuth() {
	s.Epic("Security")
	s.Feature("Authentication")
	s.Title("GET /weather/cities without key returns 401")
	s.Description("Protected weather endpoints must reject requests missing X-Internal-Key.")

	resp, err := s.weatherNoAuth.GetCities()
	s.Require().NoError(err)
	s.Assert().Equal(http.StatusUnauthorized, resp.Status,
		"expected 401, got %d — endpoint may be unprotected!", resp.Status)
}

func (s *AuthSuite) TestWeatherDataRequiresAuth() {
	s.Epic("Security")
	s.Feature("Authentication")
	s.Title("GET /weather/{city} without key returns 401")

	resp, err := s.weatherNoAuth.GetWeather("Moscow")
	s.Require().NoError(err)
	s.Assert().Equal(http.StatusUnauthorized, resp.Status,
		"expected 401, got %d", resp.Status)
}

func (s *AuthSuite) TestWeatherSyncRequiresAuth() {
	s.Epic("Security")
	s.Feature("Authentication")
	s.Title("POST /sync without key returns 401")

	resp, err := s.weatherNoAuth.TriggerSync()
	s.Require().NoError(err)
	s.Assert().Equal(http.StatusUnauthorized, resp.Status,
		"expected 401, got %d", resp.Status)
}

func (s *AuthSuite) TestAirQualityRequiresAuth() {
	s.Epic("Security")
	s.Feature("Authentication")
	s.Title("GET /air without key returns 401")

	resp, err := s.weatherNoAuth.GetAllAirQuality()
	s.Require().NoError(err)
	s.Assert().Equal(http.StatusUnauthorized, resp.Status,
		"expected 401, got %d", resp.Status)
}

// ── Races API ──────────────────────────────────────────────────────────────

func (s *AuthSuite) TestRacesListRequiresAuth() {
	s.Epic("Security")
	s.Feature("Authentication")
	s.Title("GET /races without key returns 401")
	s.Description("Protected races endpoints must reject requests missing X-Internal-Key.")

	resp, err := s.racesNoAuth.ListRaces("", false, 0)
	s.Require().NoError(err)
	s.Assert().Equal(http.StatusUnauthorized, resp.Status,
		"expected 401, got %d — endpoint may be unprotected!", resp.Status)
}

func (s *AuthSuite) TestRacesSyncRequiresAuth() {
	s.Epic("Security")
	s.Feature("Authentication")
	s.Title("POST /races/sync without key returns 401")

	resp, err := s.racesNoAuth.TriggerSync()
	s.Require().NoError(err)
	s.Assert().Equal(http.StatusUnauthorized, resp.Status,
		"expected 401, got %d", resp.Status)
}

// ── Public endpoints (should NOT require auth) ────────────────────────────

func (s *AuthSuite) TestWeatherHealthIsPublic() {
	s.Epic("Security")
	s.Feature("Authentication")
	s.Title("GET /health does not require auth on weather-api")
	s.Description("Health endpoint must be reachable without credentials for load-balancer probes.")

	resp, err := s.weatherNoAuth.GetHealth()
	s.Require().NoError(err)
	s.Assert().Equal(http.StatusOK, resp.Status,
		"health endpoint should be public, got %d", resp.Status)
}

func (s *AuthSuite) TestRacesHealthIsPublic() {
	s.Epic("Security")
	s.Feature("Authentication")
	s.Title("GET /health does not require auth on races-api")

	resp, err := s.racesNoAuth.GetHealth()
	s.Require().NoError(err)
	s.Assert().Equal(http.StatusOK, resp.Status,
		"health endpoint should be public, got %d", resp.Status)
}
