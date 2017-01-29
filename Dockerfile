FROM golang:1.6
MAINTAINER Octoblu, Inc. <docker@octoblu.com>

WORKDIR /go/src/github.com/octoblu/meshchain-util
COPY . /go/src/github.com/octoblu/meshchain-util

RUN env CGO_ENABLED=0 go build -o meshchain-util -a -ldflags '-s' .

CMD ["./meshchain-util"]
