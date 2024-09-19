FROM golang:1.23 AS builder
LABEL authors="masoud"

ENV GIN_MODE=release
ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /app

ADD v1.4.3.tar.gz .

RUN mkdir -p /usr/local/go/bin/src/github.com/gabriel-vasile/mimetype

RUN mv ./mimetype-1.4.3/* /usr/local/go/bin/src/github.com/gabriel-vasile/mimetype

RUN cd /usr/local/go/bin/src/github.com/gabriel-vasile/mimetype && go install

COPY go.mod go.sum ./

RUN echo 'replace github.com/gabriel-vasile/mimetype => /usr/local/go/bin/src/github.com/gabriel-vasile/mimetype' >> go.mod

RUN go mod tidy

COPY . .

RUN go build -o ocserv_api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ocserv_api .

COPY --from=builder /app/scripts/server.sh /server.sh

RUN chmod +x /app/ocserv_api /server.sh

ENTRYPOINT ["/server.sh"]