package dotfiles

import (
	"errors"
	"github.com/arpanrec/netcli/internal/logger"
	"github.com/go-git/go-billy/v5/osfs"
	gogit "github.com/go-git/go-git/v5"
	gogitConfig "github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"os"
)

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

	logger.Info("Adding worktree to repository: ", workTreeDir)
	storer := repository.Storer

	logger.Info("Creating worktree: ", workTreeDir)
	wt := osfs.New(workTreeDir)
	repoWt, repoWtErr := gogit.Open(storer, wt)
	if repoWtErr != nil {
		logger.Fatal("Failed to open repository with workTree: ", repoWtErr)
	}
	repository = repoWt
	logger.Info("Repository replaced with new repo with worktree")

	currentConfig, errCurrentConfig := storer.Config()
	if errCurrentConfig != nil {
		logger.Fatal("Failed to get current config: ", errCurrentConfig)
	}
	if !currentConfig.Core.IsBare {
		logger.Fatal("Repository is not bare")
	} else {
		logger.Debug("Repository is bare")
	}

	logger.Info("Setting the repository config")
	currentConfig.Core.Worktree = workTreeDir
	currentConfig.Raw.Section("status").AddOption("showUntrackedFiles", "no")
	currentConfig.Remotes["origin"] = &gogitConfig.RemoteConfig{
		Name:  "origin",
		URLs:  []string{repositoryUrl},
		Fetch: []gogitConfig.RefSpec{"+refs/heads/*:refs/remotes/origin/*"},
	}
	errConfig := repository.Storer.SetConfig(currentConfig)
	if errConfig != nil {
		logger.Fatal("Failed to set config: ", errConfig)
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
}
