#!/bin/bash
set -e

goimports -l -w $(find . -type f -name '*.go' -not -path "*/vendor/*" -not -path "*/.go/*")
gofmt -l -s -w $(find . -type f -name '*.go' -not -path "*/vendor/*" -not -path "*/.go/*")
golines -w $(find . -type f -name '*.go' -not -path "*/vendor/*" -not -path "*/.go/*")
wsl --fix ./...
golangci-lint run