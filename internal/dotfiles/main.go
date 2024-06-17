package dotfiles

import (
	"path"
	"strconv"

	"github.com/arpanrec/netcli/internal/logger"
	"github.com/arpanrec/netcli/internal/vars"
	"github.com/spf13/cobra"
)

func Main(cmd *cobra.Command, _ []string, isBackup bool) {

	workTreeDir = vars.GetHomeDir()
	backupDirRoot = path.Join(workTreeDir, ".dotfiles-backups")
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
	if isBackup {
		backupDirProvided = cmd.Flag("backup-dir").Changed
	}
	isResetHeadProvided = cmd.Flag("reset-head").Changed
	isCleanInstallProvided = cmd.Flag("clean-install").Changed

	logger.Debug("Install called with silent: ", isSilent)
	logger.Debug("Repository from flag: ", RepositoryUrl)
	logger.Debug("Branch from flag: ", Branch)
	logger.Debug("Git Directory from flag: ", GitDirectory)
	logger.Debug("Clean install flag: ", IsCleanInstall)
	logger.Debug("Reset HEAD flag: ", IsResetHead)

	preChecks()
	validateDirectoryAndLoadRepo()
	validateRepositoryUrl()
	createRemoteAuth()
	readUserInputBranch()
	install()
	if isBackup {
		backup()
	}
	checkout()
	addToRc()
	resetHead()
}
