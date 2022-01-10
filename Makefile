# This Makefile saves some typing and groups some common commands
#
# The fancy help is from here:
# https://gist.github.com/rcmachado/af3db315e31383502660#file-makefile
#
.SILENT:
.PHONY: help

## This help screen
help:
	printf "Available targets:\n\n"
	awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "%-15s %s\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)


## Build diffuser and inject the git tag and build time to variables in main
build:
	go build -ldflags "-X 'main.gitTag=$$(git describe --tags)' \
                       -X 'main.buildTimestamp=$$(date -u)'" \
                       -o diffuser_$$(go env GOOS)_$$(go env GOARCH)

## Install diffuser into /usr/local/bin (avoid standard go install)
install: build
	mv diffuser_$$(go env GOOS)_$$(go env GOARCH) \
       /usr/local/bin/diffuser_$$(go env GOOS)_$$(go env GOARCH)
	rm /usr/local/bin/diffuser
	ln -s /usr/local/bin/diffuser_$$(go env GOOS)_$$(go env GOARCH) \
       /usr/local/bin/diffuser

## Run only the unit tests
unit-tests: build
	go test ./pkg/... -v

## Run only the end to end tests
e2e-tests: install
	go test ./tests/... -v

## Run all the tests
test: unit-tests e2e-tests

## Dockerize diffuser
docker-build:
	docker build --build-arg -t evanxg852000/diffuser:latest .

## Push diffuser to DockerHub (requires DockerHub account)
docker-push: docker-build
	docker login -u evanxg852000
	docker push evanxg852000/diffuser:lates
