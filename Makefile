ELASTIC_SEARCH_IMG=dock.quicket.co.kr/elasticsearch:snp
KIBANA_IMG=dock.quicket.co.kr/kibana:snp
TEST_COMPOSE=docker-compose.test.yml
COMPOSE=docker-compose.yml

install:
	glide install

es: 
	docker build -t $(ELASTIC_SEARCH_IMG) ./build/elasticsearch/

kbn:
	docker build -t $(KIBANA_IMG) ./build/kibana/

build: es kbn

deploy:
	docker push $(ELASTIC_SEARCH_IMG) 
	docker push $(KIBANA_IMG) 

run_test:
	docker-compose -f $(TEST_COMPOSE) up -d

stop_test:
	docker-compose -f $(TEST_COMPOSE) down

run:
	docker-compose -f $(COMPOSE) up -d

stop:
	docker-compose -f $(COMPOSE) down
