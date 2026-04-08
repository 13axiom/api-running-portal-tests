# api-running-portal-tests

Integration test suite for **Running Portal** — covers `weather-api` and `races-api` REST endpoints.

**Stack:** Java 21 · Gradle 8.7 · TestNG · REST Assured · Allure

---

## Project structure

```
src/
  main/java/com/runningportal/
    config/   Config.java          — reads env vars (API URLs, key)
    model/    *.java               — POJO response models
    client/   *ApiClient.java      — REST Assured HTTP clients
    steps/    *Steps.java          — Allure @Step service layer

  test/java/com/runningportal/
    weather/  *Test.java           — weather-api test cases
    races/    *Test.java           — races-api test cases
    auth/     *Test.java           — auth guard tests (both APIs)

  test/resources/
    testng.xml                     — suite definition
```

---

## Local setup

### Prerequisites

- **Java 21** — e.g. [Temurin](https://adoptium.net/)
- **Gradle 8.7** — install via [sdkman](https://sdkman.io/):
  ```bash
  sdk install gradle 8.7
  ```
  Or download from https://gradle.org/releases/ and add to PATH.
- Both APIs running locally (see their respective READMEs)

### Environment variables

```bash
export WEATHER_API_URL=http://localhost:8080
export RACES_API_URL=http://localhost:8081
export INTERNAL_API_KEY=your_shared_secret
```

### Run all tests

```bash
gradle test
```

### Run a specific suite

```bash
gradle test -Dsuite=weather   # weather-api tests only
gradle test -Dsuite=races     # races-api tests only
gradle test -Dsuite=auth      # auth guard tests only
```

### View Allure report locally

```bash
# Install Allure CLI first: https://allurereport.org/docs/install/
allure serve build/allure-results
```

---

## CI integration

The tests are pulled by both backend repos in their GitHub Actions pipeline:

- `weather-api` CI — runs **weather suite** after unit tests pass
- `races-api` CI — runs **races suite** after unit tests pass
- This repo's own CI — runs **all suites** against both live backends

Each CI job uploads an **Allure HTML report** as a downloadable artifact.

### GitHub setup (one-time)

1. **Settings → Pages** → Source: **GitHub Actions**
2. **Settings → Secrets → Actions**, add:
   - `INTERNAL_API_KEY` — shared key matching the backend `.env`
3. If backend repos are private, also add:
   - `GH_PAT` — personal access token with `repo` scope, and uncomment the `token:` line in `ci.yml`

### Enable / disable integration tests

Integration tests run on every PR to `main` by default.
To disable temporarily, set a **Repository Variable**:

```
Settings → Secrets and variables → Actions → Variables
ENABLE_API_TESTS = false
```

Remove the variable (or set to `true`) to re-enable.

---

## Adding new tests

1. Add a test class under `src/test/java/com/runningportal/<suite>/`
2. Annotate with `@Epic`, `@Feature`, `@Story`, `@Severity` from Allure
3. Use the `*Steps` class — never call the HTTP client directly from tests
4. To add a new API call: client first → step method second → test calls step

The pattern: **Arrange → Act → Assert** with `Allure.step("...", () -> {...})` for assertion blocks.
