# syntax=docker/dockerfile:1.4

FROM golang:1.23.5 AS go

WORKDIR /api

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# RUN go install http@latest
# RUN go install github.com/gorilla/mux

# COPY . /api

RUN go build -o .

EXPOSE 6000