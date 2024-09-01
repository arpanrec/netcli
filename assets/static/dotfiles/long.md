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
