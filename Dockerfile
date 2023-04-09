# syntax=docker/dockerfile:1

FROM golang:1.20

# Set destination for COPY
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
# Copy app files
COPY . .

# Expose port
EXPOSE 8080
# Start app
CMD go run .
