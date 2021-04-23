FROM golang:1.16.3-alpine3.13 AS builder

ARG VERSION="$(git describe --tags)"
ARG COMMIT_ID="$(git rev-parse HEAD)"

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o umock -ldflags="\
-X 'github.com/chonla/umock/main.Version=${VERSION}' \
-X 'github.com/chonla/umock/main.CommitID=${COMMIT_ID}'" main.go

FROM alpine:3.13

WORKDIR /app

COPY --from=builder /app/umock .