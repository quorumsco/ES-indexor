FROM golang
MAINTAINER Wilmot Guillaume - Quorums

RUN go get github.com/tools/godep

ADD . /go/src/github.com/quorumsco/ES-indexor

WORKDIR /go/src/github.com/quorumsco/ES-indexor

RUN godep go build

ENTRYPOINT ["./ES-indexor"]
