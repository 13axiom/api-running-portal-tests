package com.runningportal.model;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

/**
 * Matches WeatherResponse from weather-api:
 * {"city":{...}, "current":{...}, "history":[...]}
 */
@JsonIgnoreProperties(ignoreUnknown = true)
public class WeatherResponse {

    public CityResponse city;
    public WeatherSnapshot current;

    @JsonIgnoreProperties(ignoreUnknown = true)
    public static class WeatherSnapshot {
        public Double  temperature;
        public Double  windspeed;
        public Double  precipitation;
        public Integer weatherCode;
        public String  recordedAt;
    }

    @Override
    public String toString() {
        return "WeatherResponse{city=" + city + ", current=" + current + '}';
    }
}
