.PHONY: setup test test-weather test-races test-auth report open-report clean

## Install dependencies
setup:
	go mod tidy

## Run ALL integration tests (requires both backends running)
test:
	go test ./tests/... -v -count=1

## Run only weather-api tests
test-weather:
	go test ./tests/weather/... -v -count=1

## Run only races-api tests
test-races:
	go test ./tests/races/... -v -count=1

## Run only auth / security tests
test-auth:
	go test ./tests/auth/... -v -count=1

## Generate Allure report from results in ./allure-results
report:
	allure generate allure-results --clean -o allure-report

## Open the report in the default browser
open-report:
	allure open allure-report

## Run tests AND immediately open the report
test-report: test report open-report

## Delete generated artefacts
clean:
	rm -rf allure-results allure-report
