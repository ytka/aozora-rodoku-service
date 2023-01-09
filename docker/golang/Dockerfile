FROM golang:1.19-alpine3.17

RUN apk update && \
    apk upgrade && \
    apk add git

RUN go install github.com/cespare/reflex@latest
ENV CGO_ENABLED=0

WORKDIR /go/src/app
COPY go.mod go.sum main.go ./

RUN go mod download