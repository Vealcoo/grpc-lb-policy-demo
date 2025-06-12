#!/bin/bash
set -e

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/main.go
chmod +x main
docker build -f Dockerfilelocal -t "demo-client:latest" .
rm main
