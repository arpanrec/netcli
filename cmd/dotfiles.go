package cmd

import (
	"github.com/arpanrec/netcli/assets"
	"github.com/arpanrec/netcli/internal/dotfiles"
	"github.com/arpanrec/netcli/internal/vars"
	"github.com/spf13/cobra"
)

func getDotFilesCmd() *cobra.Command {

	var dotFilesCmd = &cobra.Command{
		Use: "dotfiles",
		Example: `# Install dotfiles from repository
netcli dotfiles -r https://github.com/arpanrec/dotfiles.git -b main -d "${HOME}/.dotfiles"

# Install in silent mode
netcli dotfiles -r https://github.com/arpanrec/dotfiles.git -b main -d "${HOME}/.dotfiles" -s`,
		Short: "Install dotfiles",
		Long:  assets.GetTextFromTextTemplate("static/dotfiles/long.md", "dotfiles_long", nil),
		Run: func(cmd *cobra.Command, args []string) {
			dotfiles.Main(cmd, args, false)
		},
	}

	var dotFilesBackupCmd = &cobra.Command{
		Use:   "backup",
		Short: "Backup existing dotfiles",
		Long:  "Backup existing dotfiles before installing new ones.",
		Args:  vars.IDontAllowArguments,
		Example: `# Backup existing dotfiles
netcli dotfiles backup

# Backup in silent mode
netcli dotfiles -r https://github.com/arpanrec/dotfiles.git -b main -d "${HOME}/.dotfiles" -s backup`,
		Run: func(cmd *cobra.Command, args []string) {
			dotfiles.Main(cmd, args, true)
		},
	}

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

	return dotFilesCmd
}
