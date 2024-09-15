FROM golang:1.22 AS builder
LABEL authors="masoud"

ENV GIN_MODE=release
ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ocserv_api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ocserv_api .

COPY --from=builder /app/scripts/server.sh /server.sh

RUN chmod +x /app/ocserv_api /server.sh

ENTRYPOINT ["/server.sh"]