.PHONY: build start stop logs logs-tail

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
