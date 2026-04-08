package com.runningportal.races;

import com.runningportal.steps.RacesSteps;
import io.qameta.allure.*;
import io.restassured.response.Response;
import org.testng.annotations.Test;

import static org.testng.Assert.*;

@Epic("Races API")
@Feature("Sync Endpoint")
public class RacesInternalTest {

    private final RacesSteps steps = new RacesSteps();

    @Test(description = "POST /api/v1/races/sync with valid API key returns 200 or 202")
    @Story("Manual sync trigger")
    @Severity(SeverityLevel.NORMAL)
    public void syncWithValidKeySucceeds() {
        // Arrange — API key provided via X-Internal-Key header

        // Act
        Response response = steps.syncRaces();

        // Assert
        Allure.step("Status code is 200 or 202", () ->
                assertTrue(response.statusCode() == 200 || response.statusCode() == 202,
                        "Unexpected status: " + response.statusCode()));
    }

    @Test(description = "POST /api/v1/races/sync without API key returns 401 or 403")
    @Story("Sync auth guard")
    @Severity(SeverityLevel.CRITICAL)
    public void syncWithoutKeyIsRejected() {
        // Arrange — intentionally omitting X-Internal-Key header

        // Act
        Response response = steps.syncRacesWithoutKey();

        // Assert
        Allure.step("Status code is 401 or 403", () ->
                assertTrue(response.statusCode() == 401 || response.statusCode() == 403,
                        "Expected 401/403, got: " + response.statusCode()));
    }
}
