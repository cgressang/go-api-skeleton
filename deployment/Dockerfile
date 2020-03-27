FROM golang:1.13.6-alpine3.11

RUN apk update && apk upgrade && \
	apk add --no-cache ca-certificates \
	git \
	bash \
	make

RUN go get github.com/codegangsta/gin

WORKDIR $GOPATH