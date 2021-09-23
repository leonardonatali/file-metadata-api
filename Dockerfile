FROM golang:1.17-alpine3.14

WORKDIR /app

ENV GO111MODULE=off
RUN apk update; apk add git alpine-sdk
RUN go get github.com/oxequa/realize 
ENV GO111MODULE=on

COPY . /app/

EXPOSE ${PORT}
