.PHONY: install format data clear

NAME = slack-weather-message

SHELL := bash

install:
	pip install --upgrade pip \
	pip install -r requirements.txt \
	pre-commit install

format:
	isort --settings-file=isort.cfg weather.py
	flake8 --config=.flake8 weather.py

run:
	python weather.py

clear:
	rm -fr **/__pycache__
