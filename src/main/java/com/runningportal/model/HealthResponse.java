package com.runningportal.model;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

/**
 * /health response body (both APIs share the same shape).
 */
@JsonIgnoreProperties(ignoreUnknown = true)
public class HealthResponse {

    public String status;

    @Override
    public String toString() {
        return "HealthResponse{status='" + status + "'}";
    }
}
