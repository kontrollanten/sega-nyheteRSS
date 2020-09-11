FROM golang:1.15-alpine3.12

USER root
RUN apk add --update make dep

ADD . src/sega-nyheter/

WORKDIR src/sega-nyheter

RUN make && make build

ENTRYPOINT bin/fetch-feeds
