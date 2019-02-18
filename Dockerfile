FROM golang:alpine AS base

	ENV CGO_ENABLED=0

	RUN apk add --update git

	ADD ./ /go/src/github.com/cirocosta/concourse-worker-health-checker
	WORKDIR /go/src/github.com/cirocosta/concourse-worker-health-checker

	RUN go get -v
	RUN go test -v ./...
	RUN go build -v -a -tags netgo -ldflags '-w' -o /usr/local/bin/checker


FROM alpine

	COPY --from=base /usr/local/bin/checker /usr/local/bin/checker

