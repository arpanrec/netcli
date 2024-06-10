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
	repositoryUrlProvided = cmd.Flag("repositoryUrl").Changed
	branchProvided = cmd.Flag("branch").Changed
	directoryProvided = cmd.Flag("gitDirectory").Changed
	sshKeyPathProvided = cmd.Flag("ssh-key").Changed
	sshKeyPassphraseProvided = cmd.Flag("ssh-passphrase").Changed
	backupDirProvided = cmd.Flag("backupDir").Changed

	logger.Debug("Install called with silent: ", isSilent)
	logger.Debug("Repository from flag: ", repositoryUrl)
	logger.Debug("Branch from flag: ", branch)
	logger.Debug("Git Directory from flag: ", gitDirectory)
	logger.Debug("Clean install flag: ", isCleanInstall)
	logger.Debug("Reset HEAD flag: ", isResetHead)

	preChecks()
	readUserInputDirectory()
	validateDirectoryAndLoadRepo()
	readUserInputRepositoryUrl()
	validateRepositoryUrl()
	createRemoteAuth()
	readUserInputBranch()
	install()
	checkoutWithCmd()
	if cmd.Use == backupCmdUse {
		backup()
	}
}
