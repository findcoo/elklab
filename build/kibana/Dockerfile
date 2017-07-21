FROM docker.elastic.co/kibana/kibana:5.5.0
MAINTAINER usjeong <usjeong@quicket.co.kr>

COPY kibana.yml config/

USER root
RUN chown kibana:kibana config/kibana.yml

USER kibana
