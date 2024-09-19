FROM golang:1.22 AS builder
LABEL authors="masoud"

ENV GIN_MODE=release
ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /app
COPY go.mod go.sum ./

RUN mkdir /tmp/mimetype-1.4.3

COPY /root/mimetype-1.4.3 /tmp/mimetype-1.4.3

RUN mkdir -p /usr/local/go/bin/src/github.com/gabriel-vasile/mimetype

RUN mv ./mimetype-1.4.3/* /usr/local/go/bin/src/github.com/gabriel-vasile/mimetype

RUN cd /usr/local/go/bin/src/github.com/gabriel-vasile/mimetype && go install

RUN go mod download

COPY . .

RUN go build -o ocserv_api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ocserv_api .

COPY --from=builder /app/scripts/server.sh /server.sh

RUN chmod +x /app/ocserv_api /server.sh

ENTRYPOINT ["/server.sh"]