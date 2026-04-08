package com.runningportal.model;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

/**
 * Subset of the weather-api /air-quality response body.
 */
@JsonIgnoreProperties(ignoreUnknown = true)
public class AirQualityResponse {

    public String city;
    public Integer aqi;
    public String category;

    @Override
    public String toString() {
        return "AirQualityResponse{city='" + city + "', aqi=" + aqi +
               ", category='" + category + "'}";
    }
}
