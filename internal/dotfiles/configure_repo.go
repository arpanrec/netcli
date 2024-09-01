package dotfiles

import (
	"errors"
	"os"

	"github.com/arpanrec/netcli/internal/logger"
	"github.com/go-git/go-billy/v5/osfs"
	gogit "github.com/go-git/go-git/v5"
	gogitconfig "github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
)

func install() {
	logger.Info("Installing dotfiles")
	if repository == nil {
		logger.Info("Bare Cloning repository: ", RepositoryUrl)
		r, err := gogit.PlainClone(GitDirectory, true, &gogit.CloneOptions{
			URL:           RepositoryUrl,
			Auth:          authMethod,
			Progress:      os.Stdout,
			ReferenceName: plumbing.ReferenceName(Branch),
			SingleBranch:  false,
			RemoteName:    "origin",
			Tags:          gogit.AllTags,
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
	currentConfig.Core.IsBare = true
	currentConfig.Core.RepositoryFormatVersion = "0"
	showUntrackedFiles := currentConfig.Raw.Section("status").Option("showUntrackedFiles")
	if showUntrackedFiles != "no" {
		currentConfig.Raw.Section("status").SetOption("showUntrackedFiles", "no")
	}
	fileMode := currentConfig.Raw.Section("core").Option("fileMode")
	if fileMode != "true" {
		currentConfig.Raw.Section("core").SetOption("fileMode", "true")
	}
	currentConfig.Remotes["origin"] = &gogitconfig.RemoteConfig{
		Name:  "origin",
		URLs:  []string{RepositoryUrl},
		Fetch: []gogitconfig.RefSpec{"+refs/heads/*:refs/remotes/origin/*"},
	}
	errConfig := repository.Storer.SetConfig(currentConfig)
	if errConfig != nil {
		logger.Fatal("Failed to set config: ", errConfig)
	}

	logger.Info("Fetching repository: ", RepositoryUrl)
	errFetch := repository.Fetch(&gogit.FetchOptions{
		Auth:     authMethod,
		Progress: os.Stdout,
		RefSpecs: []gogitconfig.RefSpec{"+refs/heads/*:refs/remotes/origin/*"},
		Prune:    true,
	})
	if errFetch != nil {
		if errors.Is(errFetch, gogit.NoErrAlreadyUpToDate) {
			logger.Info("Repository is already up to date")
		} else {
			logger.Fatal("Failed to fetch repository: ", errFetch)
		}
	} else {
		logger.Info("Repository fetched successfully")
	}
}
