package com.runningportal.model;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * Matches Race struct from races-api:
 * {"id":1,"title":"...","race_date":"...","location":"...","region":"spb","country":"...","url":"..."}
 */
@JsonIgnoreProperties(ignoreUnknown = true)
public class RaceResponse {

    public Long   id;
    public String title;

    @JsonProperty("race_date")
    public String raceDate;

    @JsonProperty("end_date")
    public String endDate;

    public String location;
    public String region;
    public String country;
    public String distances;
    public String url;
    public String source;

    @Override
    public String toString() {
        return "RaceResponse{id=" + id + ", title='" + title +
               "', location='" + location + "', raceDate='" + raceDate + "'}";
    }
}
