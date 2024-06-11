## netcli dotfiles

Set of utilities for bootstrapping a new machine. Install dotfiles.

### Synopsis

NetCLI is a set of utilities for bootstrapping a new machine.
Setup home directory with dotfiles and configurations.

```
netcli dotfiles [flags]
```

### Options

```
  -b, --branch string           Branch to clone dotfiles from repository url, default is from ls-remote
  -c, --clean-install           Clean install, remove existing dotfiles
  -d, --git-directory string    Directory to clone dotfiles to
  -h, --help                    help for dotfiles
  -r, --repository-url string   Repository to clone dotfiles from
  -x, --reset-head              Reset HEAD to the latest commit
  -k, --ssh-key string          Path to ssh key
  -p, --ssh-passphrase string   Passphrase for ssh key
```

### Options inherited from parent commands

```
      --debug-logging   Enable debug logging
  -s, --silent          Silent mode
```

### SEE ALSO

* [netcli](netcli.md)	 - Set of utilities for bootstrapping a new machine.
* [netcli dotfiles backup](netcli_dotfiles_backup.md)	 - Set of utilities for bootstrapping a new machine. Install dotfiles. Backup existing dotfiles.

