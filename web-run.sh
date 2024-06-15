#!/usr/bin/env bash
set -euo pipefail
bash <(curl -sSL https://github.com/arpanrec/netcli/releases/download/DEVEL/netcli-DEVEL.sh) "${@}"
