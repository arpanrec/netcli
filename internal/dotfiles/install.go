package dotfiles

import (
	"github.com/arpanrec/netcli/internal/logger"
	"github.com/spf13/cobra"
	"strconv"
)

func install(cmd *cobra.Command, _ []string) {
	isS, err := strconv.ParseBool(cmd.Flag("silent").Value.String())
	if err != nil {
		logger.Fatal("Failed to get silent flag", err)
	}
	isSilent = isS
	logger.Debug("Install called with silent: ", isSilent)
	logger.Debug("Repository from flag: ", repositoryUrl)
	logger.Debug("Branch from flag: ", branch)
	logger.Debug("Directory from flag: ", directory)
	logger.Debug("Clean install flag: ", isCleanInstall)
	logger.Debug("Reset HEAD flag: ", isResetHead)

	if isSilent {
		if repositoryUrl == "" {
			logger.Fatal("Repository is not provided, but running in silent mode")
		}
		if directory == "" {
			logger.Fatal("Directory is not provided, but running in silent mode")
		}
	}

	readUserInputDirectory()
	validateDirectoryAndLoadRepo()
	readUserInputRepositoryUrl()
	validateRepositoryUrl()
	createRemoteAuth()

	// readUserInputBranch()
	// logger.Info("Branch: ", branch)

}
