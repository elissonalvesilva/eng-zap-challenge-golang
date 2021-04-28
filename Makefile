.PHONY: build run-api run-routine

CONTAINER_NAME := olx-application

build:
	docker-compose build

run-api:
	docker-compose up -d

run-routine:
	docker-compose exec CONTAINER_NAME -- /app/routine

