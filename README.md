# Netcli

Few utilities for bootstrapping a new machine

NetCLI is a set of utilities for my day-to-day work.

This helps simplify the process of setting up a new machine, installing the necessary tools, and configuring them, etc. etc.

## Installation

```bash
curl -L -o ~/.local/bin/netcli "https://github.com/arpanrec/netcli/releases/download/1.0.3/netcli-1.0.3-$(uname -s)-$(uname -m)"
chmod +x ~/.local/bin/netcli
```

From web run:

```bash
bash <(curl -s https://raw.githubusercontent.com/arpanrec/netcli/main/web-run.sh)
```

With args

```bash
bash <(curl -s https://raw.githubusercontent.com/arpanrec/netcli/main/web-run.sh) --version
```

## [Usage](docs/netcli.md)
