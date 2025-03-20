# syntax=docker/dockerfile:1.4

# Golang image as builder
FROM golang:1.24-alpine 
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

# Use ubuntu image for container
#FROM ubuntu:latest

# Set working directory
#WORKDIR /root

# Copy executable
#COPY --from=builder /api/rooms-api .

# Expose application port
EXPOSE 6000
#EXPOSE 3306

# Run application
CMD ["/rooms-api"]