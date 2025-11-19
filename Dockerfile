# ---- Build stage ----
FROM golang:1.25-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Install git and SSL certs
RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Copy dependency files and download modules
COPY go.mod go.sum ./
RUN go mod download

# Copy application code
COPY . .

# Build the server binary
RUN go build -ldflags="-s -w" -o /server ./cmd/server/main.go

# ---- Production image ----
FROM alpine:latest

RUN apk add --no-cache ca-certificates

ENV SERVER_PORT=:6969 \
    SWAPI_BASE_URL=https://swapi.dev/api \
    SWAPI_PAGE_SIZE=15

# Copy built server binary from builder
COPY --from=builder /server /server

EXPOSE 6969

ENTRYPOINT [ "/server" ]