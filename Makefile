DOCKER_COMPOSE          := docker-compose
.PHONY: test-unit
test-unit:
	go test -race ./...


.PHONY: docker-up
docker-up:
	#$(info) 'Launching the development environment'
	$(DOCKER_COMPOSE) up --build -d

.PHONY: docker-sh
docker-sh:
	$(DOCKER_COMPOSE) exec server bash

.PHONY: docker-down
docker-down:
	$(DOCKER_COMPOSE) down --remove-orphans

.PHONY: docker-logs
docker-logs:
	$(DOCKER_COMPOSE) logs -f server