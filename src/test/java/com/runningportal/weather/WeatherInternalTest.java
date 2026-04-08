package com.runningportal.weather;

import com.runningportal.steps.WeatherSteps;
import io.qameta.allure.*;
import io.restassured.response.Response;
import org.testng.annotations.Test;

import static org.testng.Assert.*;

@Epic("Weather API")
@Feature("Sync Endpoints")
public class WeatherInternalTest {

    private final WeatherSteps steps = new WeatherSteps();

    @Test(description = "POST /api/v1/sync returns 200 or 202")
    @Story("Manual weather sync")
    @Severity(SeverityLevel.NORMAL)
    public void weatherSyncSucceeds() {
        // Arrange — no preconditions needed

        // Act
        Response response = steps.syncWeather();

        // Assert
        Allure.step("Status code is 200 or 202", () ->
                assertTrue(response.statusCode() == 200 || response.statusCode() == 202,
                        "Unexpected status: " + response.statusCode()));
    }

    @Test(description = "POST /api/v1/air/sync without API key returns 401 or 403")
    @Story("Air sync auth guard")
    @Severity(SeverityLevel.CRITICAL)
    public void airSyncWithoutKeyIsRejected() {
        // Arrange — intentionally omitting X-Internal-Key header

        // Act
        Response response = steps.syncAirQualityWithoutKey();

        // Assert
        Allure.step("Status code is 401 or 403", () ->
                assertTrue(response.statusCode() == 401 || response.statusCode() == 403,
                        "Expected 401/403, got: " + response.statusCode()));
    }
}
