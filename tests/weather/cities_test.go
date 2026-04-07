package weather_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/13axiom/api-running-portal-tests/internal/client"
	"github.com/13axiom/api-running-portal-tests/internal/config"
	"github.com/joho/godotenv"
)

// ── Suite ──────────────────────────────────────────────────────────────────

type CitiesSuite struct {
	suite.Suite
	cfg *config.Config
	api *client.WeatherClient
}

func (s *CitiesSuite) BeforeAll() {
	_ = godotenv.Load("../../.env")
	s.cfg = config.Load()
	s.api = client.NewWeatherClient(s.cfg.WeatherAPIURL, s.cfg.InternalAPIKey, s.cfg.RequestTimeout)
}

func TestCities(t *testing.T) {
	suite.RunSuite(t, new(CitiesSuite))
}

// ── Test cases ─────────────────────────────────────────────────────────────

func (s *CitiesSuite) TestHealthEndpointReturns200() {
	s.Epic("Weather API")
	s.Feature("Health Check")
	s.Title("GET /health returns 200 OK")
	s.Description("Service health endpoint must be reachable without authentication.")

	resp, err := s.api.GetHealth()
	s.Require().NoError(err, "health request should not fail at transport level")
	s.Assert().Equal(http.StatusOK, resp.Status, "expected 200, got %d — is weather-api running?", resp.Status)
}

func (s *CitiesSuite) TestGetCitiesReturnsNonEmptyList() {
	s.Epic("Weather API")
	s.Feature("Cities")
	s.Title("GET /cities returns a non-empty JSON array")
	s.Description("Cities list must have at least one entry and each entry must have a name field.")

	resp, err := s.api.GetCities()
	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, resp.Status, "unexpected status: %d body: %s", resp.Status, resp.Body)

	var cities []client.City
	s.Require().NoError(resp.Parse(&cities), "response must be a valid JSON array")
	s.Assert().NotEmpty(cities, "cities list must not be empty")
}

func (s *CitiesSuite) TestEachCityHasRequiredFields() {
	s.Epic("Weather API")
	s.Feature("Cities")
	s.Title("Each city object has name, latitude, longitude")
	s.Description("All three fields are required for rendering the map and fetching weather.")

	resp, err := s.api.GetCities()
	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, resp.Status)

	var cities []client.City
	s.Require().NoError(resp.Parse(&cities))

	for _, city := range cities {
		s.Assert().NotEmpty(city.Name,      "city.name must not be empty")
		s.Assert().NotZero(city.Latitude,   "city.latitude must not be zero for %q", city.Name)
		s.Assert().NotZero(city.Longitude,  "city.longitude must not be zero for %q", city.Name)
	}
}

func (s *CitiesSuite) TestGetCitiesContentTypeIsJSON() {
	s.Epic("Weather API")
	s.Feature("Cities")
	s.Title("GET /cities response body is valid JSON")
	s.Description("Response must be parseable as JSON — not HTML error page.")

	resp, err := s.api.GetCities()
	s.Require().NoError(err)

	var raw json.RawMessage
	s.Assert().NoError(json.Unmarshal(resp.Body, &raw), "body must be valid JSON, got: %s", resp.Body)
}
