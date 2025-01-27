# syntax=docker/dockerfile:1.4

FROM golang as go

WORKDIR /api

COPY . /api/src

EXPOSE 6000