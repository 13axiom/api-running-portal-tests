package com.runningportal.weather;

import com.runningportal.model.AirQualityResponse;
import com.runningportal.steps.WeatherSteps;
import io.qameta.allure.*;
import io.restassured.response.Response;
import org.testng.annotations.DataProvider;
import org.testng.annotations.Test;

import static org.testng.Assert.*;

@Epic("Weather API")
@Feature("Air Quality Endpoint")
public class AirQualityEndpointTest {

    private final WeatherSteps steps = new WeatherSteps();

    @DataProvider(name = "cities")
    public Object[][] cities() {
        return new Object[][]{
                {"Saint Petersburg"},
                {"Limassol"},
        };
    }

    @Test(dataProvider = "cities",
          description = "GET /api/v1/air/{city} with API key returns 200, 404 or 503")
    @Story("Air quality data")
    @Severity(SeverityLevel.CRITICAL)
    public void airQualityWithApiKey(String city) {
        // Arrange
        // API key is provided via X-Internal-Key header (set in BaseClient.withApiKey())

        // Act
        Response response = steps.getAirQuality(city);

        // Assert
        Allure.step("Status code is 200, 404 or 503", () ->
                assertTrue(
                        response.statusCode() == 200 ||
                        response.statusCode() == 404 ||   // not synced yet
                        response.statusCode() == 503,     // OWM key not set on server
                        "Unexpected status: " + response.statusCode()));

        if (response.statusCode() == 200) {
            AirQualityResponse body = steps.asAirQuality(response);

            Allure.step("Body contains city name", () ->
                    assertNotNull(body.city));

            Allure.step("AQI value is in range [0, 500]", () -> {
                assertNotNull(body.aqi);
                assertTrue(body.aqi >= 0 && body.aqi <= 500,
                        "AQI out of range: " + body.aqi);
            });
        }
    }
}
