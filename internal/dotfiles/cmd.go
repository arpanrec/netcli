package dotfiles

import (
	"github.com/spf13/cobra"
	"path"
)

var Cmd = &cobra.Command{
	Use:   cmdUse,
	Short: "Install dotfiles",
	Long: `Setup home directory with dotfiles and configurations.

This command will clone the dotfiles repository and install the dotfiles in the home directory.
Git bare directory is ` + "`${HOME}/.dotfiles`." + `

The alias ` + "`dotfiles`" + ` is used to interact with the repository.

` + "```bash" + `
alias dotfiles = 'git --git-dir="${HOME}/.dotfiles" --work-tree=${HOME}'
` + "```" + `

Also all the untracked files are ignored by default.

` + "```bash" + `
dotfiles config --local status.showUntrackedFiles no
` + "```" + `

FYI: If any directory name is matching with any branch then it will cause an error. For example,
if you have a directory named ` + "`main`" + ` and you are trying to-checkout ` + "`main`" + ` branch then it will cause an error.
`,
	Run: main,
}

var dotFilesBackupCmd = &cobra.Command{
	Use:   backupCmdUse,
	Short: "Backup existing dotfiles",
	Long:  "Backup existing dotfiles before installing new ones.",
	Run:   main,
}

func init() {
	// wd, wdErr := os.UserHomeDir()
	// if wdErr != nil {
	// 	logger.Fatal("Failed to get home gitDirectory: ", wdErr)
	// }
	// workTreeDir = wd
	workTreeDir = "/home/arpan/.tmp/dotfiles_test"
	backupDirRoot = path.Join(workTreeDir, ".dotfiles-backups")

	Cmd.PersistentFlags().StringVarP(&repositoryUrl, "repository-url", "r", "",
		"Repository to clone dotfiles from")
	Cmd.PersistentFlags().StringVarP(&branch, "branch", "b", "",
		"Branch to clone dotfiles from repository url, default is from ls-remote")
	Cmd.PersistentFlags().StringVarP(&gitDirectory, "git-directory", "d", "",
		"Directory to clone dotfiles to")
	Cmd.PersistentFlags().BoolVarP(&isCleanInstall, "clean-install", "c", false,
		"Clean install, remove existing dotfiles")
	Cmd.PersistentFlags().BoolVarP(&isResetHead, "reset-head", "x", false,
		"Reset HEAD to the latest commit")
	Cmd.PersistentFlags().StringVarP(&sshKeyPath, "ssh-key", "k", "", "Path to ssh key")
	Cmd.PersistentFlags().StringVarP(&sshKeyPassphrase, "ssh-passphrase", "p", "",
		"Passphrase for ssh key")

	Cmd.AddCommand(dotFilesBackupCmd)
	dotFilesBackupCmd.PersistentFlags().StringVarP(&backupDir, "backup-dir", "u", "",
		`Directory to backup existing dotfiles. In silent mode Default: "${HOME}/.dotfiles-backups/dd-mm-yyyy"`)
}
