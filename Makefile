# Go parameterss
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=ctxm
BINARY_UNIX=$(BINARY_NAME)_unix

.PHONY: test

all: build test clean
build: 
		$(GOBUILD) -o $(BINARY_NAME) -v
test: build
		$(GOTEST) -v ./...
clean: 
		$(GOCLEAN)
		rm -f $(BINARY_NAME)