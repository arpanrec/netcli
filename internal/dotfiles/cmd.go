package dotfiles

import (
	"github.com/arpanrec/netcli/internal/constants"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "dotfiles",
	Short: constants.NetCliShort + " Install dotfiles.",
	Long:  constants.NetCliLong + "\nSetup home directory with dotfiles and configurations.",
	Run:   main,
}

var dotFilesBackupCmd = &cobra.Command{
	Use:   "backup",
	Short: Cmd.Short + " Backup existing dotfiles.",
	Long:  Cmd.Long + "\nBackup existing dotfiles before installing new ones.",
	Run:   main,
}

func init() {
	// wd, wdErr := os.UserHomeDir()
	// if wdErr != nil {
	// 	logger.Fatal("Failed to get home gitDirectory: ", wdErr)
	// }
	// workTreeDir = wd
	workTreeDir = "/home/arpan/.tmp/dotfiles_test/"

	Cmd.PersistentFlags().StringVarP(&repositoryUrl, "repositoryUrl", "r", "",
		"Repository to clone dotfiles from")
	Cmd.PersistentFlags().StringVarP(&branch, "branch", "b", "",
		"Branch to clone dotfiles from repositoryUrl, default is from ls-remote")
	Cmd.PersistentFlags().StringVarP(&gitDirectory, "gitDirectory", "d", "",
		"Directory to clone dotfiles to")
	Cmd.PersistentFlags().BoolVarP(&isCleanInstall, "clean", "c", false,
		"Clean install, remove existing dotfiles")
	Cmd.PersistentFlags().BoolVarP(&isResetHead, "reset", "R", false,
		"Reset HEAD to the latest commit")
	Cmd.PersistentFlags().StringVarP(&sshKeyPath, "ssh-key", "k", "", "Path to ssh key")
	Cmd.PersistentFlags().StringVarP(&sshKeyPassphrase, "ssh-passphrase", "p", "",
		"Passphrase for ssh key")
	Cmd.AddCommand(dotFilesBackupCmd)
}
