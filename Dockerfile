# syntax=docker/dockerfile:1
# Build the application from source
FROM golang:1-bookworm AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /crawl cmd/crawl/main.go \
    && go build -o /server cmd/server/main.go \
    && go build -o /client cmd/client/main.go

# # Run the tests in the container
# FROM build-stage AS run-test-stage
# RUN go test -v ./...

# Deploy the application binary into a lean image
FROM golang:1-bookworm AS build-release-stage

WORKDIR /

COPY --from=build-stage /crawl /usr/local/bin/
COPY --from=build-stage /server /usr/local/bin/
COPY --from=build-stage /client /usr/local/bin/
COPY ./entrypoint.sh /usr/local/bin/

RUN chmod +x /usr/local/bin/entrypoint.sh

ENTRYPOINT ["entrypoint.sh"]