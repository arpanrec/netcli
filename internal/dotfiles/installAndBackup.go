package dotfiles

import (
	"errors"
	"github.com/go-git/go-git/v5/plumbing"
	"os"
	"strconv"

	"github.com/arpanrec/netcli/internal/logger"
	gogit "github.com/go-git/go-git/v5"
	gogitConfig "github.com/go-git/go-git/v5/config"
	"github.com/spf13/cobra"
)

func installAndBackup(cmd *cobra.Command, _ []string) {
	isS, err := strconv.ParseBool(cmd.Flag("silent").Value.String())
	if err != nil {
		logger.Fatal("Failed to get silent flag", err)
	}
	isSilent = isS
	repositoryUrlProvided = cmd.Flag("repositoryUrl").Changed
	branchProvided = cmd.Flag("branch").Changed
	directoryProvided = cmd.Flag("directory").Changed
	sshKeyPathProvided = cmd.Flag("ssh-key").Changed
	sshKeyPassphraseProvided = cmd.Flag("ssh-passphrase").Changed

	logger.Debug("Install called with silent: ", isSilent)
	logger.Debug("Repository from flag: ", repositoryUrl)
	logger.Debug("Branch from flag: ", branch)
	logger.Debug("Directory from flag: ", directory)
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
}

func install() {
	if repository == nil {
		logger.Info("Bare Cloning repository: ", repositoryUrl)
		r, err := gogit.PlainClone(directory, true, &gogit.CloneOptions{
			URL:           repositoryUrl,
			Auth:          authMethod,
			Progress:      os.Stdout,
			ReferenceName: plumbing.ReferenceName(branch),
		})
		if err != nil {
			logger.Fatal("Failed to clone repository: ", err)
		}
		repository = r
	}
	if repository == nil {
		logger.Fatal("Failed to clone/open repository")
	}

	errFetch := repository.Fetch(&gogit.FetchOptions{
		Auth:     authMethod,
		Progress: os.Stdout,
		RefSpecs: []gogitConfig.RefSpec{"refs/heads/*:refs/remotes/origin/*"},
	})
	if errFetch != nil {
		if errors.Is(errFetch, gogit.NoErrAlreadyUpToDate) {
			logger.Info("Repository is already up to date")
		} else {
			logger.Fatal("Failed to fetch repository: ", errFetch)
		}
	}
}
