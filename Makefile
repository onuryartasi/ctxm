# Go parameterss
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=ctxm
BINARY_UNIX=$(BINARY_NAME)_unix
COVERAGE_NAME=coverage.out

.PHONY: test

all: 
build: 
		$(GOBUILD) -o $(BINARY_NAME) -v
test: 
		$(GOTEST) -v  -coverpkg ./... ./...  -coverprofile=$(COVERAGE_NAME)
clean: 
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(COVERAGE_NAME)
coverage: test
		$(GOCMD) tool cover -func=coverage.out