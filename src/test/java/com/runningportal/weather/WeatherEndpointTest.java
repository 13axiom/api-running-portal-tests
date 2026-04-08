package com.runningportal.weather;

import com.runningportal.model.WeatherResponse;
import com.runningportal.steps.WeatherSteps;
import io.qameta.allure.*;
import io.restassured.response.Response;
import org.testng.annotations.DataProvider;
import org.testng.annotations.Test;

import static org.testng.Assert.*;

@Epic("Weather API")
@Feature("Weather Endpoint")
public class WeatherEndpointTest {

    private final WeatherSteps steps = new WeatherSteps();

    @DataProvider(name = "knownCities")
    public Object[][] knownCities() {
        return new Object[][]{
                {"Saint Petersburg"},
                {"Limassol"},
        };
    }

    @Test(dataProvider = "knownCities",
          description = "GET /api/v1/weather/{city} returns 200 with valid body (or 404 if not synced yet)")
    @Story("Get weather for city")
    @Severity(SeverityLevel.CRITICAL)
    public void weatherForKnownCity(String city) {
        // Arrange
        // city is passed as parameter; no additional setup needed

        // Act
        Response response = steps.getWeather(city);

        // Assert
        Allure.step("Status code is 200 or 404 (404 = data not synced yet)", () ->
                assertTrue(response.statusCode() == 200 || response.statusCode() == 404,
                        "Unexpected status: " + response.statusCode()));

        if (response.statusCode() == 200) {
            WeatherResponse body = steps.asWeather(response);

            Allure.step("Body contains city object", () ->
                    assertNotNull(body.city));

            Allure.step("City name is not blank", () ->
                    assertFalse(body.city.name.isBlank()));
        }
    }

    @Test(description = "GET /api/v1/weather/{city} for unknown city returns 404")
    @Story("Unknown city handling")
    @Severity(SeverityLevel.NORMAL)
    public void unknownCityReturns404() {
        // Arrange
        String unknownCity = "XYZUnknownCity999";

        // Act
        Response response = steps.getWeather(unknownCity);

        // Assert
        Allure.step("Status code is 404", () ->
                assertEquals(response.statusCode(), 404));
    }
}
