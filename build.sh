#!/bin/bash

echo "Generate static"
go run ./cmd/genpages/gen_pages.go
echo "Build..."
GOOS=linux GOARCH=amd64 go build -v -ldflags='-s -w' -o ./dist/torrs ./cmd/main