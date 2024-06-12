package cmd

import (
	"os"
	"path"

	"github.com/arpanrec/netcli/internal/constants"
	"github.com/arpanrec/netcli/internal/dotfiles"

	"github.com/arpanrec/netcli/internal/logger"
	"github.com/spf13/cobra"
)

var dotFilesCmd = &cobra.Command{
	Use: "dotfiles",
	Example: `# Install dotfiles from repository
netcli dotfiles -r https://github.com/arpanrec/dotfiles.git -b main -d "${HOME}/.dotfiles"

# Install in silent mode
netcli dotfiles -r https://github.com/arpanrec/dotfiles.git -b main -d "${HOME}/.dotfiles" -s`,
	Short: "Install dotfiles",
	Long: `Setup home directory with dotfiles and configurations.

This command will clone the dotfiles repository and install the dotfiles in the home directory.
Git bare directory is ` + "`${HOME}/.dotfiles`." + `

The alias ` + "`dotfiles`" + ` is used to interact with the repository.

` + "```bash" + `
alias dotfiles = 'git --git-dir="${HOME}/.dotfiles" --work-tree=${HOME}'
` + "```" + `

Also, all the untracked files are ignored by default.
` + "```bash" + `
dotfiles config --local status.showUntrackedFiles no
` + "```" + `

FYI: If any directory name is matching with any branch then it will cause an error. For example,` +
		`if you have a directory named ` + "`main`" + ` and you are trying to-checkout ` + "`main`" +
		` branch then it will cause an error.

[More Details](https://wiki.archlinux.org/title/Dotfiles)`,
	Run: dotfiles.Main,
}

var dotFilesBackupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup existing dotfiles",
	Long:  "Backup existing dotfiles before installing new ones.",
	Args:  constants.IDontAllowArguments,
	Example: `# Backup existing dotfiles
netcli dotfiles backup

# Backup in silent mode
netcli dotfiles -r https://github.com/arpanrec/dotfiles.git -b main -d "${HOME}/.dotfiles" -s backup`,
	Run: dotfiles.Main,
}

func init() {
	wd, wdErr := os.UserHomeDir()
	if wdErr != nil {
		logger.Fatal("Failed to get home gitDirectory: ", wdErr)
	}
	dotfiles.WorkTreeDir = wd
	dotfiles.BackupDirRoot = path.Join(dotfiles.WorkTreeDir, ".dotfiles-backups")

	dotFilesCmd.PersistentFlags().StringVarP(&dotfiles.RepositoryUrl, "repository-url", "r", "",
		"Repository to clone dotfiles from.")
	dotFilesCmd.PersistentFlags().StringVarP(&dotfiles.Branch, "branch", "b", "",
		"Branch to clone dotfiles from repository url, default is from ls-remote if not provided and not in silent mode.")
	dotFilesCmd.PersistentFlags().StringVarP(&dotfiles.GitDirectory, "git-directory", "d", "",
		"Directory to clone dotfiles to. Default: ${HOME}/.dotfiles if not provided and not in silent mode.")
	dotFilesCmd.PersistentFlags().BoolVarP(&dotfiles.IsCleanInstall, "clean-install", "c", false,
		"Clean install, remove existing dotfiles.")
	dotFilesCmd.PersistentFlags().BoolVarP(&dotfiles.IsResetHead, "reset-head", "x", false,
		"Reset HEAD to the latest commit.")
	dotFilesCmd.PersistentFlags().StringVarP(&dotfiles.SshKeyPath, "ssh-key", "k", "",
		"Path to ssh key.")
	dotFilesCmd.PersistentFlags().StringVarP(&dotfiles.SshKeyPassphrase, "ssh-passphrase", "p", "",
		"Passphrase for ssh key.")

	dotFilesCmd.AddCommand(dotFilesBackupCmd)
	dotFilesBackupCmd.PersistentFlags().StringVarP(&dotfiles.BackupDir, "backup-dir", "u", "",
		`Directory to backup existing dotfiles. In silent mode Default: "${HOME}/.dotfiles-backups/<Unix epoch time>".`)

	netCLI.AddCommand(dotFilesCmd)
}
