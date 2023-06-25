.PHONY: install format test clear

NAME = slack-weather-message

SHELL := bash

install:
	pip install --upgrade pip \
	pip install -r requirements.txt \
	pre-commit install

format:
	isort --settings-file=isort.cfg weather.py
	flake8 --config=.flake8 weather.py

test:
	python weather.py

clear:
	shopt -s globstar ; \
	rm -fr **/__pycache__ **/*.json ;
