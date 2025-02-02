# syntax=docker/dockerfile:1.4

FROM golang:latest as go

WORKDIR /api

RUN go get -u http
RUN go get -u github.com/gorilla/mux

COPY . /api

RUN go build .

EXPOSE 6000