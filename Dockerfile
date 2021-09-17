FROM golang:1.17-alpine3.14

WORKDIR /app

ENV GO111MODULE=off
RUN apk add git
RUN go get github.com/oxequa/realize 

COPY . /app/

CMD [ "realize", "start" ]
