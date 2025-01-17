# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=cybertron
COMMIT := $(shell git rev-parse HEAD)
VERSION := "local-dev"

all: lint test
build:
	$(GOBUILD) -o $(BINARY_NAME) -ldflags "-X main.version=$(VERSION) -X main.commit=$(COMMIT) -s -w" -v
lint:
	golangci-lint run
reportcard:
	goreportcard-cli -t 100 -v
test:
	$(GOCMD) test -v --bench --benchmem -coverprofile coverage.txt -covermode=atomic ./...
	$(GOCMD) vet ./... 2> govet-report.out
	$(GOCMD) tool cover -html=coverage.txt -o cover-report.html
	printf "\nCoverage report available at cover-report.html\n\n"
theme-test-build:
	go run main.go  build -d test_data/ -o test_data/ -n "Test Theme"
tidy:
	$(GOCMD) mod tidy
clean:
	$(GOCLEAN)
	$(GOCMD) fmt ./...
	rm -f $(BINARY_NAME)
# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -ldflags "-X main.version=$(VERSION) -X main.commit=$(COMMIT) -s -w" -v
