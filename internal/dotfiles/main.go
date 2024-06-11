package dotfiles

import (
	"strconv"

	"github.com/arpanrec/netcli/internal/logger"
	"github.com/spf13/cobra"
)

func main(cmd *cobra.Command, _ []string) {
	isS, err := strconv.ParseBool(cmd.Flag("silent").Value.String())
	if err != nil {
		logger.Fatal("Failed to get silent flag", err)
	}
	isSilent = isS
	repositoryUrlProvided = cmd.Flag("repository-url").Changed
	branchProvided = cmd.Flag("branch").Changed
	gitDirectoryProvided = cmd.Flag("git-directory").Changed
	sshKeyPathProvided = cmd.Flag("ssh-key").Changed
	sshKeyPassphraseProvided = cmd.Flag("ssh-passphrase").Changed
	backupDirProvided = cmd.Flag("backup-dir").Changed
	isResetHeadProvided = cmd.Flag("reset-head").Changed
	isCleanInstallProvided = cmd.Flag("clean-install").Changed

	logger.Debug("Install called with silent: ", isSilent)
	logger.Debug("Repository from flag: ", repositoryUrl)
	logger.Debug("Branch from flag: ", branch)
	logger.Debug("Git Directory from flag: ", gitDirectory)
	logger.Debug("Clean install flag: ", isCleanInstall)
	logger.Debug("Reset HEAD flag: ", isResetHead)

	preChecks()
	validateDirectoryAndLoadRepo()
	validateRepositoryUrl()
	createRemoteAuth()
	readUserInputBranch()
	install()
	if cmd.Use == backupCmdUse {
		backup()
	}
	checkout()
}
