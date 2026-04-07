package weather_test

import (
	"net/http"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/13axiom/api-running-portal-tests/internal/client"
	"github.com/13axiom/api-running-portal-tests/internal/config"
	"github.com/joho/godotenv"
)

type WeatherSuite struct {
	suite.Suite
	cfg *config.Config
	api *client.WeatherClient
}

func (s *WeatherSuite) BeforeAll() {
	_ = godotenv.Load("../../.env")
	s.cfg = config.Load()
	s.api = client.NewWeatherClient(s.cfg.WeatherAPIURL, s.cfg.InternalAPIKey, s.cfg.RequestTimeout)
}

func TestWeather(t *testing.T) {
	suite.RunSuite(t, new(WeatherSuite))
}

// ── Test cases ─────────────────────────────────────────────────────────────

func (s *WeatherSuite) TestGetWeatherForKnownCity() {
	s.Epic("Weather API")
	s.Feature("Weather")
	s.Title("GET /weather/Moscow returns 200 with current weather")
	s.Description("Moscow is always in the seeded cities list; weather data must be present.")

	resp, err := s.api.GetWeather("Moscow")
	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, resp.Status, "body: %s", resp.Body)

	var w client.WeatherResponse
	s.Require().NoError(resp.Parse(&w))
	s.Assert().NotEmpty(w.City.Name, "city name must not be empty")
}

func (s *WeatherSuite) TestGetWeatherHasCurrentBlock() {
	s.Epic("Weather API")
	s.Feature("Weather")
	s.Title("Weather response includes current temperature and wind speed")

	resp, err := s.api.GetWeather("London")
	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, resp.Status, "body: %s", resp.Body)

	var w client.WeatherResponse
	s.Require().NoError(resp.Parse(&w))

	// Temperature range sanity: -100 to +100 °C
	s.Assert().Greater(w.Current.Temperature, -100.0, "temperature is unrealistically low")
	s.Assert().Less(w.Current.Temperature,     100.0, "temperature is unrealistically high")

	// Wind speed must be non-negative
	s.Assert().GreaterOrEqual(w.Current.WindSpeed, 0.0, "wind speed cannot be negative")
}

func (s *WeatherSuite) TestGetWeatherHasSyncedAt() {
	s.Epic("Weather API")
	s.Feature("Weather")
	s.Title("Weather response includes synced_at timestamp")
	s.Description("synced_at tells the frontend when the data was last refreshed.")

	resp, err := s.api.GetWeather("Tokyo")
	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, resp.Status, "body: %s", resp.Body)

	var w client.WeatherResponse
	s.Require().NoError(resp.Parse(&w))
	s.Assert().NotEmpty(w.Current.SyncedAt, "synced_at must be present in weather response")
}

func (s *WeatherSuite) TestGetWeatherUnknownCityReturns404() {
	s.Epic("Weather API")
	s.Feature("Weather")
	s.Title("GET /weather/UnknownXyzCity returns 404")
	s.Description("Requesting weather for a city not in the database must return 404, not 500.")

	resp, err := s.api.GetWeather("UnknownXyzCity12345")
	s.Require().NoError(err)
	s.Assert().Equal(http.StatusNotFound, resp.Status,
		"expected 404 for unknown city, got %d — body: %s", resp.Status, resp.Body)
}

func (s *WeatherSuite) TestAllSeededCitiesHaveWeather() {
	s.Epic("Weather API")
	s.Feature("Weather")
	s.Title("All seeded cities return weather data")
	s.Description("Each city returned by /cities must have weather available via /weather/{city}.")

	citiesResp, err := s.api.GetCities()
	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, citiesResp.Status)

	var cities []client.City
	s.Require().NoError(citiesResp.Parse(&cities))
	s.Require().NotEmpty(cities, "need at least one city to run this test")

	for _, city := range cities {
		resp, err := s.api.GetWeather(city.Name)
		s.Assert().NoError(err, "transport error for city %q", city.Name)
		s.Assert().Equal(http.StatusOK, resp.Status,
			"expected 200 for city %q, got %d", city.Name, resp.Status)
	}
}
