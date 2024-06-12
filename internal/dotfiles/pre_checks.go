package dotfiles

import "github.com/arpanrec/netcli/internal/logger"

func preChecks() {

	if isSilent {
		if RepositoryUrl == "" {
			logger.Fatal("Repository is not provided, but running in silent mode")
		}
		if GitDirectory == "" {
			logger.Fatal("Directory is not provided, but running in silent mode")
		}
	}

	if repositoryUrlProvided && RepositoryUrl == "" {
		logger.Fatal("Repository URL is empty")
	}

	if branchProvided && Branch == "" {
		logger.Fatal("Branch is empty")
	}

	if gitDirectoryProvided && GitDirectory == "" {
		logger.Fatal("Directory is empty")
	}

	if sshKeyPathProvided && SshKeyPath == "" {
		logger.Fatal("SSH key path is empty")
	}

	if backupDirProvided && BackupDir == "" {
		logger.Fatal("Backup directory is empty")
	}

}
