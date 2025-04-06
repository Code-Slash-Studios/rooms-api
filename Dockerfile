# syntax=docker/dockerfile:1.4

# Golang image as builder
FROM golang:1.23-alpine 
# AS builder

# Set working directory
WORKDIR /api

# Copy Go module files and dependencies
COPY go.mod ./
RUN go mod download

# Copy source code
COPY . /api

# Build binary executable
RUN go build -o /rooms-api .

# Expose application ports
EXPOSE 6000
#EXPOSE 8000

# Run application
CMD ["/rooms-api"]