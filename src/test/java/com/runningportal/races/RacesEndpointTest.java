package com.runningportal.races;

import com.runningportal.model.RaceResponse;
import com.runningportal.steps.RacesSteps;
import io.qameta.allure.*;
import io.restassured.response.Response;
import org.testng.annotations.Test;

import java.util.List;

import static org.testng.Assert.*;

@Epic("Races API")
@Feature("Races Endpoint")
public class RacesEndpointTest {

    private final RacesSteps steps = new RacesSteps();

    @Test(description = "GET /api/v1/races returns 200 and a list")
    @Story("Get upcoming races")
    @Severity(SeverityLevel.CRITICAL)
    public void racesEndpointResponds() {
        // Arrange — no preconditions needed

        // Act
        Response response = steps.getRaces();

        // Assert
        Allure.step("Status code is 200", () ->
                assertEquals(response.statusCode(), 200));

        List<RaceResponse> races = steps.asRaceList(response);

        Allure.step("Response is a list (not null)", () ->
                assertNotNull(races));
    }

    @Test(description = "GET /api/v1/races?limit=5 returns at most 5 items")
    @Story("Limit parameter")
    @Severity(SeverityLevel.NORMAL)
    public void limitParameterIsRespected() {
        // Arrange
        int limit = 5;

        // Act
        Response response = steps.getRaces(true, limit, null);

        // Assert
        Allure.step("Status code is 200", () ->
                assertEquals(response.statusCode(), 200));

        List<RaceResponse> races = steps.asRaceList(response);

        Allure.step("Response contains at most " + limit + " races", () ->
                assertTrue(races.size() <= limit,
                        "Expected <= " + limit + ", got: " + races.size()));
    }

    @Test(description = "GET /api/v1/races?region=spb returns only SPb races")
    @Story("Region filter")
    @Severity(SeverityLevel.NORMAL)
    public void regionFilterSpb() {
        // Arrange
        String region = "spb";

        // Act
        Response response = steps.getRaces(false, 50, region);

        // Assert
        Allure.step("Status code is 200", () ->
                assertEquals(response.statusCode(), 200));

        List<RaceResponse> races = steps.asRaceList(response);

        Allure.step("All returned races have region = 'spb'", () -> {
            for (RaceResponse race : races) {
                assertEquals(race.region, region,
                        "Wrong region for race: " + race.title);
            }
        });
    }

    @Test(description = "GET /api/v1/races?region=cyprus returns only Cyprus races")
    @Story("Region filter")
    @Severity(SeverityLevel.NORMAL)
    public void regionFilterCyprus() {
        // Arrange
        String region = "cyprus";

        // Act
        Response response = steps.getRaces(false, 50, region);

        // Assert
        Allure.step("Status code is 200", () ->
                assertEquals(response.statusCode(), 200));

        List<RaceResponse> races = steps.asRaceList(response);

        Allure.step("All returned races have region = 'cyprus'", () -> {
            for (RaceResponse race : races) {
                assertEquals(race.region, region,
                        "Wrong region for race: " + race.title);
            }
        });
    }

    @Test(description = "GET /api/v1/races?upcoming=false returns >= items than upcoming=true")
    @Story("Include past events")
    @Severity(SeverityLevel.NORMAL)
    public void allRacesIncludesPastEvents() {
        // Arrange — no preconditions needed

        // Act
        Response allResponse      = steps.getAllRaces(200);
        Response upcomingResponse = steps.getRaces();

        // Assert
        Allure.step("Both responses return 200", () -> {
            assertEquals(allResponse.statusCode(), 200);
            assertEquals(upcomingResponse.statusCode(), 200);
        });

        List<RaceResponse> all      = steps.asRaceList(allResponse);
        List<RaceResponse> upcoming = steps.asRaceList(upcomingResponse);

        Allure.step("All-races count >= upcoming-races count", () ->
                assertTrue(all.size() >= upcoming.size(),
                        "all=" + all.size() + " upcoming=" + upcoming.size()));
    }

    @Test(description = "Each race has id, title and raceDate fields set")
    @Story("Data completeness")
    @Severity(SeverityLevel.NORMAL)
    public void eachRaceHasRequiredFields() {
        // Arrange — no preconditions needed

        // Act
        Response response = steps.getAllRaces(100);
        List<RaceResponse> races = steps.asRaceList(response);

        // Assert
        Allure.step("Status code is 200", () ->
                assertEquals(response.statusCode(), 200));

        Allure.step("Each race has non-null id", () -> {
            for (RaceResponse race : races) {
                assertNotNull(race.id, "id is null for: " + race);
            }
        });

        Allure.step("Each race has non-null title", () -> {
            for (RaceResponse race : races) {
                assertNotNull(race.title, "title is null for: " + race);
            }
        });

        Allure.step("Each race has non-null raceDate", () -> {
            for (RaceResponse race : races) {
                assertNotNull(race.raceDate, "raceDate is null for: " + race);
            }
        });
    }
}
