package com.runningportal.auth;

import com.runningportal.steps.WeatherSteps;
import io.qameta.allure.*;
import io.restassured.response.Response;
import org.testng.annotations.Test;

import static org.testng.Assert.*;

@Epic("Auth")
@Feature("Weather API — Auth Guard")
public class AuthWeatherTest {

    private final WeatherSteps steps = new WeatherSteps();

    @Test(description = "POST /api/v1/air/sync without API key returns 401 or 403")
    @Story("Protected endpoint rejects missing key")
    @Severity(SeverityLevel.BLOCKER)
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
