package com.runningportal.weather;

import com.runningportal.model.HealthResponse;
import com.runningportal.steps.WeatherSteps;
import io.qameta.allure.*;
import io.restassured.response.Response;
import org.testng.annotations.Test;

import static org.testng.Assert.*;

@Epic("Weather API")
@Feature("Health Check")
public class WeatherApiHealthTest {

    private final WeatherSteps steps = new WeatherSteps();

    @Test(description = "GET /health returns 200 and status ok")
    @Story("Service availability")
    @Severity(SeverityLevel.BLOCKER)
    public void healthReturnsOk() {
        // Arrange — no preconditions needed

        // Act
        Response response = steps.getHealth();

        // Assert
        Allure.step("Status code is 200", () ->
                assertEquals(response.statusCode(), 200));

        HealthResponse body = steps.asHealth(response);

        Allure.step("Body field 'status' equals 'ok'", () ->
                assertEquals(body.status.toLowerCase(), "ok"));
    }
}
