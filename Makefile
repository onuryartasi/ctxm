# Go parameterss
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=ctxm
BINARY_UNIX=$(BINARY_NAME)_unix
COVERAGE_NAME=coverage.out
CONFIG_DIR=$(HOME)/.context-manager

.PHONY: test

all: 
build: 
		$(GOBUILD) -o $(BINARY_NAME) -v
test: 
		rm -rf $(CONFIG_DIR)
		$(GOTEST) -v  -coverpkg ./... ./...  -coverprofile=$(COVERAGE_NAME)
clean: 
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(COVERAGE_NAME)
coverage: test
		$(GOCMD) tool cover -func=coverage.out