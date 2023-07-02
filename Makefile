NAME = slack-weather-message

SHELL := bash

GOCMD=go
GOBUILD=$(GOCMD) build
GOMOD=$(GOCMD) mod
GOTEST=$(GOCMD) run
GOINSTALL=$(GOCMD) install
GOFLAGS := -v

init:
 	$(GOMOD) init
	$(GOINSTALL) $(GOFLAGS) github.com/PuerkitoBio/goquery@latest

test:
	$(GOTEST) main.go

tidy:
	$(GOMOD) tidy

lint:
	golangci-lint run --enable-all

clear:
	shopt -s globstar ; \
	rm -fr **/*.json ;
