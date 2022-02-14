FROM golang:1.17.3-alpine

RUN mkdir -p /usr/src/myapp
WORKDIR /usr/src/myapp
COPY .  /usr/src/myapp

RUN go mod download