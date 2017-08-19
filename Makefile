ELASTIC_SEARCH_IMG=dock.quicket.co.kr/elasticsearch:snp
KIBANA_IMG=dock.quicket.co.kr/kibana:snp
TEST_COMPOSE=docker-compose.test.yml
COMPOSE=docker-compose.yml

install: ## install dependency
	glide install

es: ## build elasticsearch image
	docker build -t $(ELASTIC_SEARCH_IMG) ./build/elasticsearch/

kbn: ## build kibana image
	docker build -t $(KIBANA_IMG) ./build/kibana/

build: es kbn ## build all image

deploy: ## run containers
	docker push $(ELASTIC_SEARCH_IMG) 
	docker push $(KIBANA_IMG) 

run.test: ## run containers with test mode
	docker-compose -f $(TEST_COMPOSE) up -d

stop.test:  ## stop containers with test mode
	docker-compose -f $(TEST_COMPOSE) down

run: ## run containers
	docker-compose -f $(COMPOSE) up -d

stop: ## stop containers
	docker-compose -f $(COMPOSE) down

.PHONY: help
.DEFAULT_GOAL := help

help:
	@grep -E '^[a-zA-Z0-9._-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
