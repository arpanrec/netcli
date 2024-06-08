package dotfiles

import (
	"errors"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/storage/filesystem"
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
	directoryProvided = cmd.Flag("gitDirectory").Changed
	sshKeyPathProvided = cmd.Flag("ssh-key").Changed
	sshKeyPassphraseProvided = cmd.Flag("ssh-passphrase").Changed

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
}

func install() {
	if repository == nil {
		logger.Info("Bare Cloning repository: ", repositoryUrl)
		r, err := gogit.PlainClone(gitDirectory, true, &gogit.CloneOptions{
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
	currentConfig, errCurrentConfig := repository.Config()
	if errCurrentConfig != nil {
		logger.Fatal("Failed to get current config: ", errCurrentConfig)
	}
	if !currentConfig.Core.IsBare {
		logger.Fatal("Repository is not bare")
	} else {
		logger.Debug("Repository is bare")
	}

	logger.Info("Fetching repository: ", repositoryUrl)
	errFetch := repository.Fetch(&gogit.FetchOptions{
		Auth:     authMethod,
		Progress: os.Stdout,
		RefSpecs: []gogitConfig.RefSpec{"+refs/heads/*:refs/remotes/origin/*"},
		Prune:    true,
	})
	if errFetch != nil {
		if errors.Is(errFetch, gogit.NoErrAlreadyUpToDate) {
			logger.Info("Repository is already up to date")
		} else {
			logger.Fatal("Failed to fetch repository: ", errFetch)
		}
	}
	homeDir, errHome := os.UserHomeDir()
	if errHome != nil {
		logger.Fatal("Failed to get home gitDirectory: ", errHome)
	}
	cfg, errStorerConfig := repository.Storer.Config()
	if errStorerConfig != nil {
		logger.Fatal("Failed to get config: ", errStorerConfig)
	}
	cfg.Core.Worktree = homeDir
	cfg.Raw.Section("status").AddOption("showUntrackedFiles", "no")
	cfg.Remotes["origin"] = &gogitConfig.RemoteConfig{
		Name:  "origin",
		URLs:  []string{repositoryUrl},
		Fetch: []gogitConfig.RefSpec{"+refs/heads/*:refs/remotes/origin/*"},
	}
	errConfig := repository.Storer.SetConfig(cfg)
	if errConfig != nil {
		logger.Fatal("Failed to set config: ", errConfig)
	}

	repositoryFs := osfs.New(gitDirectory)
	s := filesystem.NewStorage(repositoryFs, cache.NewObjectLRUDefault())
	wt := osfs.New(homeDir)
	r, errR := gogit.Open(s, wt)
	if errR != nil {
		logger.Fatal("Failed to open repository: ", errR)
	}

	workTree, errWorkTree := r.Worktree()
	if errWorkTree != nil {
		logger.Fatal("Failed to get worktree: ", errWorkTree)
	}
	logger.Info("Checking out branch: ", branch)
	errCheckout := workTree.Checkout(&gogit.CheckoutOptions{
		Branch: plumbing.ReferenceName(branch),
	})
	if errCheckout != nil {
		logger.Fatal("Failed to checkout branch: ", errCheckout)
	}
}
