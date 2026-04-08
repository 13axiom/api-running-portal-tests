package com.runningportal.steps;

import com.runningportal.client.WeatherApiClient;
import com.runningportal.model.AirQualityResponse;
import com.runningportal.model.CityResponse;
import com.runningportal.model.HealthResponse;
import com.runningportal.model.WeatherResponse;
import io.qameta.allure.Step;
import io.restassured.response.Response;

/**
 * Action layer for weather-api.
 * Each method performs exactly one HTTP call and returns the response.
 * No assertions here — assertions belong in the test.
 */
public class WeatherSteps {

    private final WeatherApiClient client = new WeatherApiClient();

    @Step("GET /health")
    public Response getHealth() {
        return client.getHealth();
    }

    @Step("GET /api/v1/cities")
    public Response getCities() {
        return client.getCities();
    }

    @Step("GET /api/v1/weather/{city}")
    public Response getWeather(String city) {
        return client.getWeather(city);
    }

    @Step("GET /api/v1/air/{city}")
    public Response getAirQuality(String city) {
        return client.getAirQuality(city);
    }

    @Step("POST /api/v1/sync")
    public Response syncWeather() {
        return client.syncWeather();
    }

    @Step("POST /api/v1/air/sync (no API key)")
    public Response syncAirQualityWithoutKey() {
        return client.syncAirQualityUnauthorized();
    }

    // ── Convenience deserializers (no assertions) ─────────────────────────────

    public HealthResponse asHealth(Response response) {
        return response.as(HealthResponse.class);
    }

    public CityResponse[] asCities(Response response) {
        return response.as(CityResponse[].class);
    }

    public WeatherResponse asWeather(Response response) {
        return response.as(WeatherResponse.class);
    }

    public AirQualityResponse asAirQuality(Response response) {
        return response.as(AirQualityResponse.class);
    }
}
