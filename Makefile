NAME = slack-weather-message

SHELL := bash

GOCMD=go
GOBUILD=$(GOCMD) build
GOMOD=$(GOCMD) mod
GOTEST=$(GOCMD) run
GOINSTALL=$(GOCMD) install
GOFLAGS := -v
LDFLAGS := -s -w

ifneq ($(shell go env GOOS),darwin)
LDFLAGS := -extldflags "-static"
endif

install:
	pre-commit install
	$(GOINSTALL) github.com/PuerkitoBio/goquery@latest

build:
	$(GOBUILD) $(GOFLAGS) -ldflags '$(LDFLAGS)' -o main.go

test:
	$(GOTEST) main.go

tidy:
	$(GOMOD) tidy

lint:
	golangci-lint run --enable-all

clear:
	shopt -s globstar ; \
	rm -fr **/*.json ;
