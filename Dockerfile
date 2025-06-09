FROM golang:1.24.4-alpine3.22 AS builder

ARG VERSION=""
ARG COMMIT_ID=""

WORKDIR /app

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -o umock -ldflags="-X 'main.Version=${VERSION}' -X 'main.CommitID=${COMMIT_ID}'" .

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/umock .