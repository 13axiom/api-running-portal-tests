package com.runningportal.config;

/**
 * Reads test configuration from environment variables.
 * <p>
 * Set these before running tests (locally via .env / export, in CI via GitHub secrets / env block):
 * <pre>
 *   WEATHER_API_URL     – base URL of weather-api,  default http://localhost:8080
 *   RACES_API_URL       – base URL of races-api,    default http://localhost:8081
 *   INTERNAL_API_KEY    – shared secret for /internal/** endpoints
 * </pre>
 */
public final class Config {

    public static final String WEATHER_API_URL;
    public static final String RACES_API_URL;
    public static final String INTERNAL_API_KEY;

    static {
        WEATHER_API_URL  = env("WEATHER_API_URL",  "http://localhost:8080");
        RACES_API_URL    = env("RACES_API_URL",    "http://localhost:8081");
        INTERNAL_API_KEY = env("INTERNAL_API_KEY", "");
    }

    private Config() {}

    private static String env(String name, String defaultValue) {
        String v = System.getenv(name);
        return (v != null && !v.isBlank()) ? v.trim() : defaultValue;
    }
}
