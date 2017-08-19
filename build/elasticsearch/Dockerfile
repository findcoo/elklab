FROM docker.elastic.co/elasticsearch/elasticsearch:5.5.0
MAINTAINER findcoo <thirdlif2@gmail.com>

COPY elasticsearch.yml config/
RUN elasticsearch-plugin install https://github.com/open-korean-text/elasticsearch-analysis-openkoreantext/releases/download/5.5.0.2/elasticsearch-analysis-openkoreantext-5.5.0.2-plugin.zip

USER root
RUN chown elasticsearch:elasticsearch config/elasticsearch.yml

USER elasticsearch
