package com.runningportal.weather;

import com.runningportal.model.CityResponse;
import com.runningportal.steps.WeatherSteps;
import io.qameta.allure.*;
import io.restassured.response.Response;
import org.testng.annotations.Test;

import static org.testng.Assert.*;

@Epic("Weather API")
@Feature("Cities Endpoint")
public class CitiesEndpointTest {

    private final WeatherSteps steps = new WeatherSteps();

    @Test(description = "GET /api/v1/cities returns 200 and non-empty list")
    @Story("Supported cities")
    @Severity(SeverityLevel.NORMAL)
    public void citiesReturnsNonEmptyList() {
        // Arrange — no preconditions needed

        // Act
        Response response = steps.getCities();

        // Assert
        Allure.step("Status code is 200", () ->
                assertEquals(response.statusCode(), 200));

        CityResponse[] cities = steps.asCities(response);

        Allure.step("List contains at least one city", () ->
                assertTrue(cities.length > 0));
    }

    @Test(description = "GET /api/v1/cities — each city has id and non-blank name")
    @Story("Supported cities")
    @Severity(SeverityLevel.NORMAL)
    public void eachCityHasRequiredFields() {
        // Arrange — no preconditions needed

        // Act
        Response response = steps.getCities();
        CityResponse[] cities = steps.asCities(response);

        // Assert
        Allure.step("Each city has non-null id", () -> {
            for (CityResponse city : cities) {
                assertNotNull(city.id, "city.id is null for: " + city);
            }
        });

        Allure.step("Each city has non-blank name", () -> {
            for (CityResponse city : cities) {
                assertNotNull(city.name,        "city.name is null for: "  + city);
                assertFalse(city.name.isBlank(), "city.name is blank for: " + city);
            }
        });
    }
}
