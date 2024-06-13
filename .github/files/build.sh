#!/usr/bin/env bash

NETCLI_VERSION="${NETCLI_VERSION:-"DEVEL"}"
go clean -cache -modcache -i -r
rm -rf build

tee internal/constants/version.go <<EOF
package constants

const Version = "${NETCLI_VERSION}"
EOF

go run ./main.go gendocs
# env GOOS= GOARCH= go build -o build/netcli-"${NETCLI_VERSION}"-$(uname -s)-$(uname -m)
env GOOS=linux GOARCH=arm64 go build -o build/netcli-"${NETCLI_VERSION}"-Linux-aarch64
env GOOS=linux GOARCH=amd64 go build -o build/netcli-"${NETCLI_VERSION}"-Linux-x86_64

cd build || exit 1

sha256sum netcli-"${NETCLI_VERSION}"-Linux-aarch64 | tee -a netcli-"${NETCLI_VERSION}".sha256
sha256sum netcli-"${NETCLI_VERSION}"-Linux-x86_64 | tee -a netcli-"${NETCLI_VERSION}".sha256
