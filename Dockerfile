FROM golang:1.8

RUN mkdir -p /go/src/blog-api

WORKDIR /go/src/blog-api

ADD . /go/src/blog-api

RUN go get
RUN go build ./main.go

CMD go run main.go
