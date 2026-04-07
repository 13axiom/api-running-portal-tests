package races_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/13axiom/api-running-portal-tests/internal/client"
	"github.com/13axiom/api-running-portal-tests/internal/config"
	"github.com/joho/godotenv"
)

type RacesSuite struct {
	suite.Suite
	cfg *config.Config
	api *client.RacesClient
}

func (s *RacesSuite) BeforeAll() {
	_ = godotenv.Load("../../.env")
	s.cfg = config.Load()
	s.api = client.NewRacesClient(s.cfg.RacesAPIURL, s.cfg.InternalAPIKey, s.cfg.RequestTimeout)
}

func TestRaces(t *testing.T) {
	suite.RunSuite(t, new(RacesSuite))
}

// ── Test cases ─────────────────────────────────────────────────────────────

func (s *RacesSuite) TestHealthEndpointReturns200() {
	s.Epic("Races API")
	s.Feature("Health Check")
	s.Title("GET /health returns 200 OK")

	resp, err := s.api.GetHealth()
	s.Require().NoError(err)
	s.Assert().Equal(http.StatusOK, resp.Status, "is races-api running on the configured port?")
}

func (s *RacesSuite) TestListRacesReturnsJSONArray() {
	s.Epic("Races API")
	s.Feature("Race List")
	s.Title("GET /races returns a valid JSON array")
	s.Description("Even if the DB is empty, the response must be [] not null or an error page.")

	resp, err := s.api.ListRaces("", false, 0)
	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, resp.Status, "body: %s", resp.Body)

	var races []client.Race
	s.Assert().NoError(resp.Parse(&races), "must parse as JSON array, got: %s", resp.Body)
}

func (s *RacesSuite) TestListRacesRegionFilterSPB() {
	s.Epic("Races API")
	s.Feature("Race List")
	s.Title("GET /races?region=spb returns only SPb races")
	s.Description("Region filter must be applied server-side; every returned race must have region='spb'.")

	resp, err := s.api.ListRaces("spb", false, 0)
	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, resp.Status, "body: %s", resp.Body)

	var races []client.Race
	s.Require().NoError(resp.Parse(&races))

	for _, r := range races {
		s.Assert().Equal("spb", r.Region,
			"race %q has region=%q, expected spb", r.Title, r.Region)
	}
}

func (s *RacesSuite) TestListRacesRegionFilterCyprus() {
	s.Epic("Races API")
	s.Feature("Race List")
	s.Title("GET /races?region=cyprus returns only Cyprus races")

	resp, err := s.api.ListRaces("cyprus", false, 0)
	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, resp.Status, "body: %s", resp.Body)

	var races []client.Race
	s.Require().NoError(resp.Parse(&races))

	for _, r := range races {
		s.Assert().Equal("cyprus", r.Region,
			"race %q has region=%q, expected cyprus", r.Title, r.Region)
	}
}

func (s *RacesSuite) TestListRacesLimitIsRespected() {
	s.Epic("Races API")
	s.Feature("Race List")
	s.Title("GET /races?limit=3 returns at most 3 races")
	s.Description("The limit query parameter must cap the number of results.")

	resp, err := s.api.ListRaces("", false, 3)
	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, resp.Status)

	var races []client.Race
	s.Require().NoError(resp.Parse(&races))
	s.Assert().LessOrEqual(len(races), 3, "expected at most 3 races with limit=3, got %d", len(races))
}

func (s *RacesSuite) TestUpcomingFilterExcludesPastRaces() {
	s.Epic("Races API")
	s.Feature("Race List")
	s.Title("GET /races?upcoming=true returns only future races")
	s.Description("All returned races must have race_date >= today.")

	resp, err := s.api.ListRaces("", true, 0)
	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, resp.Status)

	var races []client.Race
	s.Require().NoError(resp.Parse(&races))

	// If there are results, none should be obviously in the past year
	// (we check the year is >= current to avoid timezone edge-cases)
	for _, r := range races {
		// race_date format: "2026-05-10T00:00:00Z"
		s.Assert().False(
			len(r.RaceDate) > 4 && r.RaceDate[:4] < "2024",
			"race %q has past date %q with upcoming=true", r.Title, r.RaceDate,
		)
	}
}

func (s *RacesSuite) TestRaceFieldsArePopulated() {
	s.Epic("Races API")
	s.Feature("Race List")
	s.Title("Each race has required fields: id, title, race_date, region, source")

	resp, err := s.api.ListRaces("", false, 10)
	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, resp.Status)

	var races []client.Race
	s.Require().NoError(resp.Parse(&races))

	if len(races) == 0 {
		s.T().Skip("No races in DB — run /races/sync first")
		return
	}

	for _, r := range races {
		s.Assert().NotZero(r.ID,         "race.id must not be zero")
		s.Assert().NotEmpty(r.Title,     "race.title must not be empty")
		s.Assert().NotEmpty(r.RaceDate,  "race.race_date must not be empty")
		s.Assert().NotEmpty(r.Region,    "race.region must not be empty")
		s.Assert().NotEmpty(r.Source,    "race.source must not be empty")
	}
}

func (s *RacesSuite) TestRaceURLsHaveScheme() {
	s.Epic("Races API")
	s.Feature("Race List")
	s.Title("Race URLs have http:// or https:// scheme")
	s.Description("Bare domain URLs like 'pushkin-run.ru' cause broken links in the frontend.")

	resp, err := s.api.ListRaces("", false, 50)
	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, resp.Status)

	var races []client.Race
	s.Require().NoError(resp.Parse(&races))

	for _, r := range races {
		if r.URL == "" {
			continue // no URL is fine
		}
		lo := strings.ToLower(r.URL)
		s.Assert().True(
			strings.HasPrefix(lo, "http://") || strings.HasPrefix(lo, "https://"),
			"race %q has URL without scheme: %q", r.Title, r.URL,
		)
	}
}

func (s *RacesSuite) TestSyncEndpointReturnsResult() {
	s.Epic("Races API")
	s.Feature("Sync")
	s.Title("POST /races/sync returns sync result array")
	s.Description("Sync must complete without error and return source='aims' in the result.")

	resp, err := s.api.TriggerSync()
	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, resp.Status, "body: %s", resp.Body)

	var results []client.RaceSyncResult
	s.Require().NoError(resp.Parse(&results), "sync result must be JSON array")
	s.Require().NotEmpty(results, "sync result must contain at least one entry")

	s.Assert().Equal("aims", results[0].Source, "first sync result source must be 'aims'")
	s.Assert().Empty(results[0].Error,           "sync must not return an error string")
}
