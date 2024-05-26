# syntax=docker/dockerfile:1
# Build the application from source
FROM golang:1-bookworm AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -o crawl cmd/crawl/main.go \
    && CGO_ENABLED=0 go build -o server cmd/server/main.go \
    && CGO_ENABLED=1 go build -o client cmd/client/main.go

# # Run the tests in the container
# FROM build-stage AS run-test-stage
# RUN go test -v ./...

# Deploy the application binary into a lean image
FROM debian:12-slim AS build-release-stage

WORKDIR /

COPY --from=busybox:stable-uclibc /bin/sh /bin/sh

COPY --from=build-stage /app/crawl /usr/local/bin/
COPY --from=build-stage /app/server /usr/local/bin/
COPY --from=build-stage /app/client /usr/local/bin/
COPY --chmod=755 ./entrypoint.sh /usr/local/bin/

ENTRYPOINT ["entrypoint.sh"]