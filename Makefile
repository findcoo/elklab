install:
	glide install

es: 
	docker build -t dock.quicket.co.kr/elasticsearch:snp ./build/elasticsearch/

kbn:
	docker build -t dock.quicket.co.kr/kibana:snp ./build/kibana/

build: es kbn

deploy:
	docker push dock.quicket.co.kr/elasticsearch:snp 
	docker push dock.quicket.co.kr/kibana:snp

start:
	docker-compose up -d

stop:
	docker-compose down
