package com.runningportal.steps;

import com.runningportal.client.RacesApiClient;
import com.runningportal.model.HealthResponse;
import com.runningportal.model.RaceResponse;
import io.qameta.allure.Step;
import io.restassured.response.Response;

import java.util.Arrays;
import java.util.List;

/**
 * Action layer for races-api.
 * Each method performs exactly one HTTP call and returns the response.
 * No assertions here — assertions belong in the test.
 */
public class RacesSteps {

    private final RacesApiClient client = new RacesApiClient();

    @Step("GET /health")
    public Response getHealth() {
        return client.getHealth();
    }

    @Step("GET /api/v1/races")
    public Response getRaces() {
        return client.getRaces();
    }

    @Step("GET /api/v1/races?upcoming={upcoming}&limit={limit}&region={region}")
    public Response getRaces(boolean upcoming, int limit, String region) {
        return client.getRaces(upcoming, limit, region);
    }

    @Step("GET /api/v1/races?upcoming=false&limit={limit}")
    public Response getAllRaces(int limit) {
        return client.getAllRaces(limit);
    }

    @Step("POST /api/v1/races/sync (with API key)")
    public Response syncRaces() {
        return client.syncRaces();
    }

    @Step("POST /api/v1/races/sync (no API key)")
    public Response syncRacesWithoutKey() {
        return client.syncRacesUnauthorized();
    }

    // ── Convenience deserializers (no assertions) ─────────────────────────────

    public HealthResponse asHealth(Response response) {
        return response.as(HealthResponse.class);
    }

    public List<RaceResponse> asRaceList(Response response) {
        return Arrays.asList(response.as(RaceResponse[].class));
    }
}
