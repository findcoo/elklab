FROM docker.elastic.co/elasticsearch/elasticsearch:5.5.0
MAINTAINER usjeong <usjeong@quicket.co.kr>

COPY elasticsearch.yml config/

USER root
RUN chown elasticsearch:elasticsearch config/elasticsearch.yml

USER elasticsearch
