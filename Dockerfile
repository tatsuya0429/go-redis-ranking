FROM golang:1.18-alpine

WORKDIR /app

ENV GO111MODULE=on

RUN go install github.com/cosmtrek/air@latest

CMD [ "air" ]