DOCKER_COMPOSE          := docker-compose
.PHONY: test-unit
test-unit:
	go test -race ./...

.PHONY: test-functional
test-functional:
	go test -race -tags functional -v ./test/functional/... --count=1

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

.PHONY: install-migrate
install-migrate:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

.PHONY: create-migration
create-migration:
	migrate create -dir internal/infra/postgres/migrations -ext sql ${NAME}

.PHONY: run-migrations
run-migrations:
	migrate -source file://internal/infra/postgres/migrations -database postgres://postgres:admin@localhost:5432/example?sslmode=disable up
