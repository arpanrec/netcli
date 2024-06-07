package dotfiles

import (
	"github.com/arpanrec/netcli/internal/logger"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "dotfiles",
	Short: "Install dotfiles",
	Long:  `Setup home directory with dotfiles and configurations.`,
	Run:   install,
}

var dotFilesBackupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup existing dotfiles",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("Backup called with silent: ", cmd.Flag("silent").Value)
	},
}

func init() {
	Cmd.PersistentFlags().StringVarP(&repositoryUrl, "repositoryUrl", "r", "",
		"Repository to clone dotfiles from")
	Cmd.PersistentFlags().StringVarP(&branch, "branch", "b", "",
		"Branch to clone dotfiles from repositoryUrl, default is from ls-remote")
	Cmd.PersistentFlags().StringVarP(&directory, "directory", "d", "",
		"Directory to clone dotfiles to")
	Cmd.PersistentFlags().BoolVarP(&isCleanInstall, "clean", "c", false,
		"Clean install, remove existing dotfiles")
	Cmd.PersistentFlags().BoolVarP(&isResetHead, "reset", "R", false,
		"Reset HEAD to the latest commit")
	Cmd.AddCommand(dotFilesBackupCmd)
}
