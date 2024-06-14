#!/usr/bin/env bash
set -euo pipefail
bash <(curl -sSL https://github.com/arpanrec/netcli/releases/download/1.0.8/netcli-1.0.8.sh) "${@}"
