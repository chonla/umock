FROM golang:1.16.3-alpine3.13 AS builder

ARG VERSION=""
ARG COMMIT_ID=""

RUN apk update && apk add --no-cache --update git gcc g++

WORKDIR /go/src/github.com/chonla/umock

COPY . .

RUN go mod download

# RUN go build -o umock -ldflags="-X 'main.Version=${VERSION}' -X 'main.CommitID=${COMMIT_ID}'"
RUN go build -o umock

# RUN /app/umock version

# FROM alpine:3.13

# WORKDIR /app

# COPY --from=builder /app/umock .