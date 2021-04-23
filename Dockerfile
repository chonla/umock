FROM golang:1.16.3-alpine3.13 AS builder

ARG VERSION=""
ARG COMMIT_ID=""

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o umock -ldflags="\
-X 'main.Version=$(VERSION)' \
-X 'main.CommitID=$(COMMIT_ID)'" main.go

FROM alpine:3.13

WORKDIR /app

COPY --from=builder /app/umock .