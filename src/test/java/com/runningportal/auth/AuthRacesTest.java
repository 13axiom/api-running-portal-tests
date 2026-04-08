package com.runningportal.auth;

import com.runningportal.steps.RacesSteps;
import io.qameta.allure.*;
import io.restassured.response.Response;
import org.testng.annotations.Test;

import static org.testng.Assert.*;

@Epic("Auth")
@Feature("Races API — Auth Guard")
public class AuthRacesTest {

    private final RacesSteps steps = new RacesSteps();

    @Test(description = "POST /api/v1/races/sync without API key returns 401 or 403")
    @Story("Protected endpoint rejects missing key")
    @Severity(SeverityLevel.BLOCKER)
    public void racesSyncWithoutKeyIsRejected() {
        // Arrange — intentionally omitting X-Internal-Key header

        // Act
        Response response = steps.syncRacesWithoutKey();

        // Assert
        Allure.step("Status code is 401 or 403", () ->
                assertTrue(response.statusCode() == 401 || response.statusCode() == 403,
                        "Expected 401/403, got: " + response.statusCode()));
    }
}
