#!/usr/bin/env bash
go run ./main.go gendocs
env GOOS=linux GOARCH=arm64 go build -o build/netcli-linux-arm64
env GOOS=linux GOARCH=amd64 go build -o build/netcli-linux-amd64
