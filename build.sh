#!/bin/bash

# exit immediately on failure
set -e

BASEDIR=$(pwd)/$(dirname "$0")
cd "$BASEDIR"

CLI_EXE_NAME=dcos-github-bot

# ---

# go

go get
go fmt

CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build -ldflags="-s -w" -o $CLI_EXE_NAME"-darwin"
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o $CLI_EXE_NAME"-linux"
