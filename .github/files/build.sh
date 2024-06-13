#!/usr/bin/env bash

NETCLI_VERSION="${NETCLI_VERSION:-"DEVEL"}"

tee internal/constants/version.go <<EOF
package constants

const Version = "${NETCLI_VERSION}"
EOF

go run ./main.go gendocs
env GOOS=linux GOARCH=arm64 go build -o build/netcli-"${NETCLI_VERSION}"-linux-arm64
env GOOS=linux GOARCH=amd64 go build -o build/netcli-"${NETCLI_VERSION}"-linux-amd64
