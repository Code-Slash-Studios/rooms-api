# syntax=docker/dockerfile:1.4

FROM golang:1.23 AS go

WORKDIR /api

COPY go.mod go.sum ./

RUN go mod download

COPY /src/controllers /api/controllers
COPY /src/models /api/models
COPY /src/routes /api/routes

# RUN go install http@latest
# RUN go install github.com/gorilla/mux

COPY . /api

# WORKDIR /api/controllers
# RUN go build -o .
# WORKDIR /api/models
# RUN go build -o .
# WORKDIR /api/routes

COPY . /api

RUN go build -o /rooms-api

EXPOSE 6000

CMD ["/rooms-api"]