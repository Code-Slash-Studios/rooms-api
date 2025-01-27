# syntax=docker/dockerfile:1.4

FROM golang:latest as go

WORKDIR /api

COPY . /api

EXPOSE 6000