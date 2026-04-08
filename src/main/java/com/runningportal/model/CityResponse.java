package com.runningportal.model;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

/**
 * Matches City struct from weather-api:
 * {"id":1,"name":"Saint Petersburg","latitude":59.95,"longitude":30.32,"created_at":"..."}
 */
@JsonIgnoreProperties(ignoreUnknown = true)
public class CityResponse {
    public Integer id;
    public String  name;
    public Double  latitude;
    public Double  longitude;

    @Override
    public String toString() {
        return "CityResponse{id=" + id + ", name='" + name + "'}";
    }
}
