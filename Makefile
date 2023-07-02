NAME = slack-weather-message

SHELL := bash

GOCMD=go
GOMOD=$(GOCMD) mod
GORUN=$(GOCMD) run

run:
	$(GORUN) main.go

tidy:
	$(GOMOD) tidy

lint:
	golangci-lint run --enable-all

clear:
	shopt -s globstar ; \
	rm -fr **/*.json ;
