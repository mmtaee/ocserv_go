#!/bin/sh

export GIN_MODE=release

export CGO_ENABLED=0

export GOOS=linux

go build -o ocserv_api .