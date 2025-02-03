# syntax=docker/dockerfile:1.4

FROM golang:1.23.5 AS go

WORKDIR /api

COPY . .

# RUN go install http@latest
# RUN go install github.com/gorilla/mux

COPY . /api

RUN go build .

EXPOSE 6000