package dotfiles

import "github.com/arpanrec/netcli/internal/logger"

func preChecks() {

	if isSilent {
		if repositoryUrl == "" {
			logger.Fatal("Repository is not provided, but running in silent mode")
		}
		if gitDirectory == "" {
			logger.Fatal("Directory is not provided, but running in silent mode")
		}
	}

	if repositoryUrlProvided && repositoryUrl == "" {
		logger.Fatal("Repository URL is empty")
	}

	if branchProvided && branch == "" {
		logger.Fatal("Branch is empty")
	}

	if gitDirectoryProvided && gitDirectory == "" {
		logger.Fatal("Directory is empty")
	}

	if sshKeyPathProvided && sshKeyPath == "" {
		logger.Fatal("SSH key path is empty")
	}

	if backupDirProvided && backupDir == "" {
		logger.Fatal("Backup directory is empty")
	}

}
