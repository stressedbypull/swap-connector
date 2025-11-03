FROM golang:1.25-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk add --no-cache git ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -ldflags="-s -w" -o /server ./cmd/server/main.go

# Minimal image for final container (multi-stage build)

FROM alpine:latest
RUN apk add --no-cache ca-certificates

COPY --from=builder /server /server
EXPOSE 6969
ENTRYPOINT [ "/server" ]