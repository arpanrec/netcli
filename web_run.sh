#!/usr/bin/env bash
set -euo pipefail

NETCLI_VERSION_BIN_PATH=/tmp/netcli-1.0.0-$(uname -s)-$(uname -m)
NETCLI_VERSION_SHA256_FILE_PATH=/tmp/netcli-1.0.0.sha256
export NETCLI_VERSION_BIN_PATH

if [ ! -f "${NETCLI_VERSION_BIN_PATH}" ]; then
    curl -L -o "${NETCLI_VERSION_BIN_PATH}" "https://github.com/arpanrec/netcli/releases/download/1.0.0/netcli-1.0.0-$(uname -s)-$(uname -m)"
fi

if [ ! -f "${NETCLI_VERSION_SHA256_FILE_PATH}" ]; then
    curl -L -o "${NETCLI_VERSION_SHA256_FILE_PATH}" "https://github.com/arpanrec/netcli/releases/download/1.0.0/netcli-1.0.0.sha256"
fi

cd /tmp || exit 1

sha256sum -c "${NETCLI_VERSION_SHA256_FILE_PATH}" --ignore-missing --status

chmod +x "${NETCLI_VERSION_BIN_PATH}"
${NETCLI_VERSION_BIN_PATH} "${@}"
