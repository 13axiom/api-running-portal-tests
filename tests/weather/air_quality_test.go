package weather_test

import (
	"net/http"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/13axiom/api-running-portal-tests/internal/client"
	"github.com/13axiom/api-running-portal-tests/internal/config"
	"github.com/joho/godotenv"
)

type AirQualitySuite struct {
	suite.Suite
	cfg *config.Config
	api *client.WeatherClient
}

func (s *AirQualitySuite) BeforeAll() {
	_ = godotenv.Load("../../.env")
	s.cfg = config.Load()
	s.api = client.NewWeatherClient(s.cfg.WeatherAPIURL, s.cfg.InternalAPIKey, s.cfg.RequestTimeout)
}

func TestAirQuality(t *testing.T) {
	suite.RunSuite(t, new(AirQualitySuite))
}

func (s *AirQualitySuite) TestGetAllAirQualityReturnsArray() {
	s.Epic("Weather API")
	s.Feature("Air Quality")
	s.Title("GET /air returns a JSON array")
	s.Description("Air quality endpoint must return a JSON array (possibly empty if OWM key not configured).")

	resp, err := s.api.GetAllAirQuality()
	s.Require().NoError(err)
	// Accept 200 (data present) or 200 empty array — both are valid
	s.Assert().Equal(http.StatusOK, resp.Status, "body: %s", resp.Body)

	var results []client.AirQualitySnapshot
	s.Assert().NoError(resp.Parse(&results), "body must be a JSON array, got: %s", resp.Body)
}

func (s *AirQualitySuite) TestGetCityAirQualityHasAQIField() {
	s.Epic("Weather API")
	s.Feature("Air Quality")
	s.Title("GET /air/Moscow returns snapshot with aqi field")
	s.Description("AQI must be between 1 (Good) and 5 (Very Poor) per OpenWeatherMap scale.")

	resp, err := s.api.GetCityAirQuality("Moscow")
	s.Require().NoError(err)

	if resp.Status == http.StatusNotFound {
		s.T().Skip("Air quality data not yet synced for Moscow — run /air/sync first")
		return
	}

	s.Require().Equal(http.StatusOK, resp.Status, "body: %s", resp.Body)

	var snap client.AirQualitySnapshot
	s.Require().NoError(resp.Parse(&snap))
	s.Assert().GreaterOrEqual(snap.AQI, 1, "AQI must be >= 1")
	s.Assert().LessOrEqual(snap.AQI,   5, "AQI must be <= 5")
}

func (s *AirQualitySuite) TestAirQualityPM25IsNonNegative() {
	s.Epic("Weather API")
	s.Feature("Air Quality")
	s.Title("PM2.5 value in air quality snapshot is non-negative")

	resp, err := s.api.GetCityAirQuality("London")
	s.Require().NoError(err)

	if resp.Status == http.StatusNotFound {
		s.T().Skip("Air quality data not yet synced for London")
		return
	}
	s.Require().Equal(http.StatusOK, resp.Status)

	var snap client.AirQualitySnapshot
	s.Require().NoError(resp.Parse(&snap))
	s.Assert().GreaterOrEqual(snap.PM25, 0.0, "PM2.5 cannot be negative")
}
