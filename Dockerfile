# Use the official Golang image as the base
FROM golang:1.23 AS builder
LABEL authors="masoud"

# Set environment variables
ENV GIN_MODE=release
ENV CGO_ENABLED=0
ENV GOOS=linux

# Explicitly set the GOPATH to /go, which is the default for the Go image
ENV GOPATH=/go

# Set the working directory inside the container
WORKDIR /app

# Step 1: Create the necessary directory structure for the mimetype package
RUN mkdir -p /go/src/github.com/gabriel-vasile/mimetype

# Step 2: Add and extract the downloaded mimetype package (v1.4.3.tar.gz)
ADD v1.4.3.tar.gz /app/

# Step 3: Move the extracted mimetype package to the correct Go module path
RUN mv ./mimetype-1.4.3/* /go/src/github.com/gabriel-vasile/mimetype

# Step 4: Copy go.mod and go.sum into the working directory early
COPY go.mod go.sum ./

# Step 5: Add the replace directive in go.mod to point to the local mimetype package
RUN echo 'replace github.com/gabriel-vasile/mimetype => /go/src/github.com/gabriel-vasile/mimetype' >> go.mod

# Step 6: Download dependencies without fetching mimetype from the internet
RUN go mod download

COPY . .

RUN go build -o ocserv_api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ocserv_api .

COPY --from=builder /app/scripts/server.sh /server.sh

RUN chmod +x /app/ocserv_api /server.sh

ENTRYPOINT ["/server.sh"]