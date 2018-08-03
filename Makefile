.PHONY: build

URL := https://developers.strava.com/swagger/swagger.json

build:
	swagger-codegen generate -i "$(URL)" -l go -o .

