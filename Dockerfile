FROM golang:1.17.3-alpine

RUN mkdir -p /usr/src/myapp
WORKDIR /usr/src/myapp