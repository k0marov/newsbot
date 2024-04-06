# syntax=docker/dockerfile:1
FROM golang:1.21 AS build-stage
WORKDIR /app

# .env is not copied here, because env params are passed through docker compose (which reads .env for you)
COPY internal/ ./internal
COPY cmd/ ./cmd
COPY go.mod .
COPY go.sum .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/cli

FROM alpine AS build-release-stage

WORKDIR /

COPY --from=build-stage /main /main

ENTRYPOINT ["/main"]
