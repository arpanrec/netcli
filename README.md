# Netcli

Few utilities for bootstrapping a new machine

NetCLI is a set of utilities for my day-to-day work.

This helps simplify the process of setting up a new machine, installing the necessary tools, and configuring them, etc. etc.

## [Debian](/scripts/debian-cloudinit.sh)

Variables:

* `CLOUD_INIT_GROUP` - Group name for the user to be created. Default `cloudinit`.
* `CLOUD_INIT_USER` - Username for the user to be created. Default `cloudinit`.
* `CLOUD_INIT_USE_SSH_PUB` - Use SSH public key for the user.
* `CLOUD_INIT_IS_DEV_MACHINE` - Install development tools. Default `false`.
* `CLOUD_INIT_COPY_ROOT_SSH_KEYS` - Copy root SSH keys to the user. Default `false`.
* `CLOUD_INIT_HOSTNAME` - Hostname for the machine. Default `cloudinit`.
* `CLOUD_INIT_DOMAIN` - Domain name for the machine. Default `cloudinit`.

```bash
sudo -E -H -u root bash -c '/bin/bash <(curl \
    -s https://raw.githubusercontent.com/arpanrec/netcli/main/scripts/debian-cloudinit.sh)'
```

or for development machine

```bash
CLOUD_INIT_IS_DEV_MACHINE=true sudo -E -H -u root \
    bash -c '/bin/bash <(curl -s https://raw.githubusercontent.com/arpanrec/netcli/main/scripts/debian-cloudinit.sh)'
```

[Linode stack script](https://cloud.linode.com/stackscripts/1164660) example:

```bash
#!/bin/bash
# <UDF name="CLOUD_INIT_COPY_ROOT_SSH_KEYS" Label="Copy Root SSH Keys to current user" oneOf="true,false" default="true"/>
# <UDF name="CLOUD_INIT_IS_DEV_MACHINE" Label="Install Devel tool chain" oneOf="true,false" default="false"/>
sudo -E -H -u root bash -c '/bin/bash <(curl -s \
    https://raw.githubusercontent.com/arpanrec/netcli/main/scripts/debian-cloudinit.sh)'
```
