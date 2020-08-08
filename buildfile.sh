#!/bin/sh
set -e

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/ma-novel-crawler-api main.go
