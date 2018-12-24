FROM golang:alpine AS base

RUN apk add --update git

ADD ./ /go/src/github.com/cirocosta/concourse-worker-health-checker
WORKDIR /go/src/github.com/cirocosta/concourse-worker-health-checker

RUN set -x && \
	go get -v && \
	go build -v -o /usr/local/bin/checker


FROM alpine
COPY --from=base /usr/local/bin/checker /usr/local/bin/checker

