# syntax=docker/dockerfile:1.4

FROM golang AS go

WORKDIR /api

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go get -u http
RUN go get -u github.com/gorilla/mux

COPY . /api

RUN go build .

EXPOSE 6000