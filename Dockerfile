FROM golang:1.18-alpine as builder
RUN apk add --no-cache ca-certificates git
WORKDIR /app

COPY . /app/

RUN go mod download
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init

ENV CGO_ENABLED=0
RUN go build -ldflags="-s -w" -o /app/main.go -o /app/cmd

ENTRYPOINT goose -dir "./server/db/migrations" "postgres" "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" up && ./cmd
