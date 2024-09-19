FROM golang:1.23 AS builder

LABEL authors="masoud"

# Set environment variables
ENV GIN_MODE=release
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOPATH=/go

# Set working directory for the app
WORKDIR /app

# Step 1: Copy go.mod and go.sum first (helps with caching layers in Docker)
COPY go.mod go.sum ./

# Step 2: Add the replace directive in go.mod to point to the local mimetype package
RUN echo 'replace github.com/gabriel-vasile/mimetype => /go/src/github.com/gabriel-vasile/mimetype' >> go.mod

# Step 3: Create the necessary directory structure for the mimetype package
RUN mkdir -p /go/src/github.com/gabriel-vasile/mimetype

# Step 4: Add and extract the downloaded mimetype package (v1.4.3.tar.gz)
ADD v1.4.3.tar.gz /app/

# Step 5: Move the extracted mimetype package to the correct Go module path in GOPATH
RUN mv ./mimetype-1.4.3/* /go/src/github.com/gabriel-vasile/mimetype

COPY . .

# Step 6: Build without running go mod tidy or go mod download, to prevent downloading dependencies
RUN go build -o ocserv_api .



RUN go build -o ocserv_api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ocserv_api .

COPY --from=builder /app/scripts/server.sh /server.sh

RUN chmod +x /app/ocserv_api /server.sh

ENTRYPOINT ["/server.sh"]