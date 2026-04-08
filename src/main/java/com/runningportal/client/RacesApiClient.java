package com.runningportal.client;

import com.runningportal.config.Config;
import io.restassured.response.Response;

import static io.restassured.RestAssured.given;

/**
 * Low-level HTTP client for races-api.
 *
 * Real routes:
 *   GET  /health                — public
 *   GET  /api/v1/races          — requires X-Internal-Key
 *   POST /api/v1/races/sync     — requires X-Internal-Key
 */
public class RacesApiClient extends BaseClient {

    public RacesApiClient() {
        super(Config.RACES_API_URL);
    }

    // ── Health ────────────────────────────────────────────────────────────────

    public Response getHealth() {
        return given(spec).get("/health");
    }

    // ── Races (all require API key) ───────────────────────────────────────────

    public Response getRaces() {
        return given(withApiKey()).get("/api/v1/races");
    }

    public Response getRaces(boolean upcoming, int limit, String region) {
        var req = given(withApiKey())
                .queryParam("upcoming", upcoming)
                .queryParam("limit", limit);
        if (region != null && !region.isBlank()) {
            req = req.queryParam("region", region);
        }
        return req.get("/api/v1/races");
    }

    public Response getAllRaces(int limit) {
        return given(withApiKey())
                .queryParam("upcoming", false)
                .queryParam("limit", limit)
                .get("/api/v1/races");
    }

    // ── Sync ──────────────────────────────────────────────────────────────────

    public Response syncRaces() {
        return given(withApiKey()).post("/api/v1/races/sync");
    }

    public Response syncRacesUnauthorized() {
        return given(spec).post("/api/v1/races/sync");
    }
}
