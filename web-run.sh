#!/usr/bin/env bash
set -euo pipefail
bash <(curl -sSL https://github.com/arpanrec/netcli/releases/download/1.1.5/netcli-web-run-1.1.5.sh) "${@}"
