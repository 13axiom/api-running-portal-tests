package com.runningportal.client;

import com.runningportal.config.Config;
import io.restassured.response.Response;

import static io.restassured.RestAssured.given;

/**
 * Low-level HTTP client for weather-api.
 *
 * Real routes (all under /api/v1):
 *   GET  /health              — public
 *   GET  /api/v1/cities       — public
 *   GET  /api/v1/weather/{city} — public
 *   POST /api/v1/sync         — public (no key required per current impl)
 *   GET  /api/v1/air          — requires X-Internal-Key
 *   GET  /api/v1/air/{city}   — requires X-Internal-Key
 *   POST /api/v1/air/sync     — requires X-Internal-Key
 */
public class WeatherApiClient extends BaseClient {

    public WeatherApiClient() {
        super(Config.WEATHER_API_URL);
    }

    // ── Health ────────────────────────────────────────────────────────────────

    public Response getHealth() {
        return given(spec).get("/health");
    }

    // ── Weather (public) ──────────────────────────────────────────────────────

    public Response getWeather(String city) {
        return given(spec).pathParam("city", city).get("/api/v1/weather/{city}");
    }

    public Response getCities() {
        return given(spec).get("/api/v1/cities");
    }

    // ── Air quality (requires API key) ────────────────────────────────────────

    public Response getAirQuality(String city) {
        return given(withApiKey()).pathParam("city", city).get("/api/v1/air/{city}");
    }

    public Response getAllAirQuality() {
        return given(withApiKey()).get("/api/v1/air");
    }

    // ── Sync (public POST) ────────────────────────────────────────────────────

    public Response syncWeather() {
        return given(spec).post("/api/v1/sync");
    }

    // ── Air sync (requires API key) ───────────────────────────────────────────

    public Response syncAirQuality() {
        return given(withApiKey()).post("/api/v1/air/sync");
    }

    public Response syncAirQualityUnauthorized() {
        return given(spec).post("/api/v1/air/sync");
    }
}
