#!/usr/bin/env bash

if [ -z "${NETCLI_VERSION}" ]; then
    echo "NETCLI_VERSION is not set. Please set the version of the netcli you want to build"
    exit 1
fi

tee internal/constants/version.go <<EOF
package constants

const Version = "${NETCLI_VERSION}"
EOF

go run ./main.go gendocs
env GOOS=linux GOARCH=arm64 go build -o build/netcli-"${NETCLI_VERSION}"-linux-arm64
env GOOS=linux GOARCH=amd64 go build -o build/netcli-"${NETCLI_VERSION}"-linux-amd64
