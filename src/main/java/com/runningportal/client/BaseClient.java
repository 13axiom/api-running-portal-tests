package com.runningportal.client;

import com.runningportal.config.Config;
import io.qameta.allure.restassured.AllureRestAssured;
import io.restassured.builder.RequestSpecBuilder;
import io.restassured.filter.log.RequestLoggingFilter;
import io.restassured.filter.log.ResponseLoggingFilter;
import io.restassured.http.ContentType;
import io.restassured.specification.RequestSpecification;

/**
 * Common REST Assured request specification factory.
 * Attaches Allure filter so every HTTP call is recorded in the report.
 */
public abstract class BaseClient {

    protected final RequestSpecification spec;

    protected BaseClient(String baseUrl) {
        spec = new RequestSpecBuilder()
                .setBaseUri(baseUrl)
                .setContentType(ContentType.JSON)
                .addFilter(new AllureRestAssured())
                .addFilter(new RequestLoggingFilter())
                .addFilter(new ResponseLoggingFilter())
                .build();
    }

    /**
     * Returns a spec with the internal API key header pre-set.
     */
    protected RequestSpecification withApiKey() {
        return new RequestSpecBuilder()
                .addRequestSpecification(spec)
                .addHeader("X-Internal-Key", Config.INTERNAL_API_KEY)
                .build();
    }
}
