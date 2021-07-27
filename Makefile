COMPOSE=docker-compose.yml

install: ## install dependency
	glide install

run: ## run containers
	docker-compose -f $(COMPOSE) up -d

stop: ## stop containers
	docker-compose -f $(COMPOSE) down

.PHONY: help
.DEFAULT_GOAL := help

help:
	@grep -E '^[a-zA-Z0-9._-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
