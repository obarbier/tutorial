DOCKER_COMPOSE=docker-compose
GO=/usr/local/go/bin/go
# REQUIRED SECTION
ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
# END OF REQUIRED SECTION
DOCKER_COMPOSE_FILE=$(ROOT_DIR)/docker-compose.yaml
.PHONY: help dependencies up start stop restart status ps clean
c=
dependencies: check-dependencies ## Check dependencies

build: pusher-service-build meow-service-build query-service-build

pusher-service-build:
	@$(GO) build -o pusher_service github.com/obarbier/cqrs-go/pusher-service

meow-service-build:
	@$(GO) build -o meow_service github.com/obarbier/cqrs-go/meow-service

query-service-build:
	@$(GO) build -o query_service github.com/obarbier/cqrs-go/query-service

up: ## Start all or c=<name> containers in foreground
	@$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) up $(c)

start: ## Start all or c=<name> containers in background
	@$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) up -d $(c)

stop: ## Stop all or c=<name> containers
	@$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) stop $(c)

restart: ## Restart all or c=<name> containers
	@$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) stop $(c)
	@$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) up $(c) -d

logs: ## Show logs for all or c=<name> containers
	@$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) logs --tail=100 -f $(c)

status: ## Show status of containers
	@$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) ps

ps: status ## Alias of status

clean: confirm ## Clean all data
	@$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) down