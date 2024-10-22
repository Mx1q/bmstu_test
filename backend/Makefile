.PHONY: test allure report

test:
	rm -rf allure-results
	export ALLURE_OUTPUT_PATH="/Users/maximhalitov/Desktop/IU7/bmstu_test/backend" && go test ./... --race --parallel 10
	cp environment.properties allure-results

allure:
	[ -d allure-reports ] && cp -rf allure-reports/history allure-results || :
	rm -rf allure-reports
	allure generate allure-results -o allure-reports
	allure serve allure-results -p 4000

report: test allure

ci-unit:
	export ALLURE_OUTPUT_PATH="${GITHUB_WORKSPACE}" && \
	export ALLURE_OUTPUT_FOLDER="unit-allure" && \
 	go test -tags=unit_test ./... --race --parallel 10

ci-integration:
	export ALLURE_OUTPUT_PATH="${GITHUB_WORKSPACE}" && \
	export ALLURE_OUTPUT_FOLDER="integration-allure" && \
	go test -tags=integration_test ./... --race --parallel 10

ci-concat-reports:
	mkdir allure-results
	cp unit-allure/* allure-results/
	cp integration-allure/* allure-results/
	#cp e2e-allure/* allure-results/
	cp environment.properties allure-results