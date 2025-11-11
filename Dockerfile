FROM golang:1.25-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk add --no-cache git ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Install swag for generating Swagger documentation
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

# Generate Swagger docs
RUN swag init -g cmd/server/main.go -o ./docs

RUN go build -ldflags="-s -w" -o /server ./cmd/server/main.go

# Minimal image for final container (multi-stage build)

FROM alpine:latest
RUN apk add --no-cache ca-certificates

# Default environment variables (can be overridden by docker-compose or runtime)
ENV SERVER_PORT=:6969 \
    SWAPI_BASE_URL=https://swapi.dev/api \
    SWAPI_PAGE_SIZE=15

COPY --from=builder /server /server
EXPOSE 6969
ENTRYPOINT [ "/server" ]