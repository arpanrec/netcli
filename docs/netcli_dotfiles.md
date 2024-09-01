## netcli dotfiles

Install dotfiles

### Synopsis

Setup home directory with dotfiles and configurations.

This command will clone the dotfiles repository and install the dotfiles in the home directory.
Git bare directory is `${HOME}/.dotfiles`.

The alias `dotfiles` is used to interact with the repository.

```bash
alias dotfiles = 'git --git-dir="${HOME}/.dotfiles" --work-tree=${HOME}'
```

Also, all the untracked files are ignored by default.

```bash
dotfiles config --local status.showUntrackedFiles no
```

FYI: If any directory name is matching with any branch then it will cause an error.
For example, if you have a directory named `main` and you are trying to-checkout `main` branch then it will cause an error.

[More Details](https://wiki.archlinux.org/title/Dotfiles)

**Note:** Do you use Arch? ARCH BTW BTW, and you know what you are doing.

```bash
rm -rf "${HOME}/.dotfiles"
git clone --bare https://github.com/arpanrec/dotfiles.git "${HOME}/.dotfiles"
git --git-dir="${HOME}/.dotfiles" --work-tree="${HOME}" config --local status.showUntrackedFiles no
git --git-dir="${HOME}/.dotfiles" --work-tree="${HOME}" checkout main --force
git --git-dir="${HOME}/.dotfiles" --work-tree="${HOME}" config --local remote.origin.fetch "+refs/heads/*:refs/remotes/origin/*"
```


```
netcli dotfiles [flags]
```

### Examples

```
# Install dotfiles from repository
netcli dotfiles -r https://github.com/arpanrec/dotfiles.git -b main -d "${HOME}/.dotfiles"

# Install in silent mode
netcli dotfiles -r https://github.com/arpanrec/dotfiles.git -b main -d "${HOME}/.dotfiles" -s

# Clean installation from web run with reset Head
bash <(curl -s https://raw.githubusercontent.com/arpanrec/netcli/main/web-run.sh) dotfiles \
	-r https://github.com/arpanrec/dotfiles.git -b main -d "${HOME}/.dotfiles" -s --reset-head

```

### Options

```
  -b, --branch string           Branch to clone dotfiles from repository url, default is from ls-remote if not provided and not in silent mode.
  -c, --clean-install           Clean install, remove existing dotfiles.
  -d, --git-directory string    Directory to clone dotfiles to. Default: ${HOME}/.dotfiles if not provided and not in silent mode.
  -h, --help                    help for dotfiles
  -r, --repository-url string   Repository to clone dotfiles from.
  -x, --reset-head              Reset HEAD to the latest commit.
  -k, --ssh-key string          Path to ssh key.
  -p, --ssh-passphrase string   Passphrase for ssh key.
```

### Options inherited from parent commands

```
      --debug-logging   Enable debug logging. This can be set using the environment variable DEBUG=true.
  -s, --silent          Silent mode. Do not prompt for any input.
```

### SEE ALSO

* [netcli](netcli.md)	 - Few utilities for bootstrapping a new machine
* [netcli dotfiles backup](netcli_dotfiles_backup.md)	 - Backup existing dotfiles

