# syntax=docker/dockerfile:1
FROM golang:1.21 AS build-stage
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /main

FROM alpine AS build-release-stage

WORKDIR /

COPY --from=build-stage /main /main

ENTRYPOINT ["/main"]
