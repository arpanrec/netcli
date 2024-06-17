## netcli dotfiles backup

Backup existing dotfiles

### Synopsis

Backup existing dotfiles before installing new ones.

```
netcli dotfiles backup [flags]
```

### Examples

```
# Backup existing dotfiles
netcli dotfiles backup

# Backup in silent mode
netcli dotfiles -r https://github.com/arpanrec/dotfiles.git -b main -d "${HOME}/.dotfiles" -s backup
```

### Options

```
  -u, --backup-dir string   Directory to backup existing dotfiles. In silent mode Default: "${HOME}/.dotfiles-backups/<Unix epoch time>".
  -h, --help                help for backup
```

### Options inherited from parent commands

```
  -b, --branch string           Branch to clone dotfiles from repository url, default is from ls-remote if not provided and not in silent mode.
  -c, --clean-install           Clean install, remove existing dotfiles.
      --debug-logging           Enable debug logging. This can be set using the environment variable DEBUG=true.
  -d, --git-directory string    Directory to clone dotfiles to. Default: ${HOME}/.dotfiles if not provided and not in silent mode.
  -r, --repository-url string   Repository to clone dotfiles from.
  -x, --reset-head              Reset HEAD to the latest commit.
  -s, --silent                  Silent mode. Do not prompt for any input.
  -k, --ssh-key string          Path to ssh key.
  -p, --ssh-passphrase string   Passphrase for ssh key.
```

### SEE ALSO

* [netcli dotfiles](netcli_dotfiles.md)	 - Install dotfiles

