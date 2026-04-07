# api-tests

Integration test suite for **weather-api** and **races-api** backends.

Built with Go + [allure-go](https://github.com/ozontech/allure-go) — tests run with `go test`, results are published as an interactive Allure HTML report on GitHub Pages.

## Test coverage

| Suite | Tests | What's checked |
|-------|-------|----------------|
| **Cities** | 4 | `/health`, `/cities` list non-empty, required fields, valid JSON |
| **Weather** | 5 | `/weather/{city}` 200, temperature range, synced_at present, 404 for unknown city, all seeded cities |
| **Air Quality** | 3 | `/air` array, AQI 1-5, PM2.5 non-negative |
| **Races** | 8 | `/races` array, region filter SPb + Cyprus, limit, upcoming filter, required fields, URLs have scheme, sync result |
| **Auth** | 8 | Each protected endpoint returns 401 without key; health endpoints are public |

**28 tests total** across 5 suites.

## Run locally

### Prerequisites
- Both backends running (see their READMEs)
- Go 1.21+
- [Allure CLI](https://allurereport.org/docs/install-for-macos/) for HTML reports (optional)

```bash
# macOS
brew install allure
```

### Setup

```bash
git clone https://github.com/13axiom/api-running-portal-tests.git
cd api-tests
make setup          # go mod tidy

cp .env.example .env
# .env is pre-filled with localhost defaults
```

### Run tests

```bash
# All tests
make test

# Single suite
make test-weather
make test-races
make test-auth

# Run + generate + open report in one command
make test-report
```

### View the Allure report

```bash
# After running tests:
make report         # generates HTML in ./allure-report/
make open-report    # opens in browser
```

## GitHub CI

`.github/workflows/ci.yml` does the following on every push to `main` / `develop`:

1. **Starts PostgreSQL** as a service container
2. **Checks out** `weather-api` and `races-api` side-by-side
3. **Starts both backends** and waits for their `/health` endpoints
4. **Runs all tests** — failures don't stop the pipeline (so the report is always generated)
5. **Generates Allure HTML report** with run history / trend charts
6. **Publishes to GitHub Pages** — direct link is printed in the CI summary

### GitHub setup (one-time)

1. Go to repo **Settings → Pages**
   - Source: **GitHub Actions**

2. Go to **Settings → Secrets and variables → Actions**, add:
   - `INTERNAL_API_KEY` — shared key from `.env` files
   - `OWM_API_KEY` — OpenWeatherMap key (for air quality tests)

3. If your backend repos are private, also add:
   - `GH_PAT` — a personal access token with `repo` scope, then update `ci.yml` to use it

After the first successful run, the report is available at:
```
https://13axiom.github.io/api-tests/
```

## Project structure

```
api-tests/
  internal/
    config/          — loads env vars
    client/
      base.go        — shared HTTP client (do, Get, Post, Parse)
      weather_client.go — typed methods for weather-api
      races_client.go   — typed methods for races-api
  tests/
    weather/
      cities_test.go      — /health, /cities
      weather_test.go     — /weather/{city}
      air_quality_test.go — /air, /air/{city}
    races/
      races_test.go       — /races, /races/sync
    auth/
      auth_test.go        — 401 checks + public endpoint checks
  .github/workflows/ci.yml
  Makefile
  .env.example
```
