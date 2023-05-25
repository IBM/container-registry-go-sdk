# Makefile to build go-sdk-template

all: build unittest lint tidy

travis-ci: build unittest lint tidy

build:
	go build ./...

unittest:
	go test `go list ./... | grep -v samples`

alltest:
	go test `go list ./... | grep -v samples` -v -tags=integration
	go test `go list ./... | grep -v samples` -v -tags=examples

registryintegration:
	go test ./containerregistryv1 -v -tags=integration

vaintegration:
	go test ./vulnerabilityadvisorv4 -v -tags=integration

lint:
	golangci-lint run

tidy:
	go mod tidy
