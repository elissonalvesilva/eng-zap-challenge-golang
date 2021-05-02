.PHONY: build start stop logs logs-tail tests-docker tests-local

CONTAINER_NAME := olx-application

build:
	docker-compose build

start:
	docker-compose up -d

stop:
	docker-compose down

logs:
	docker logs -f $(CONTAINER_NAME) 

logs-tail:
	docker logs -f --tail 100 $(CONTAINER_NAME)

tests-docker:
	docker build --no-cache -t $(CONTAINER_NAME)-test -f ./Dockerfile.test .
	docker run -v ${PWD}:/go/testdir $(CONTAINER_NAME)-test

tests-local:
	go test -v ./tests/...