package dotfiles

import (
	"errors"
	"net/url"
	"strings"

	"github.com/arpanrec/netcli/internal/logger"
	"github.com/arpanrec/netcli/internal/utils"
	"github.com/manifoldco/promptui"
)

func readUserInputRepositoryUrl() {
	if RepositoryUrl != "" {
		return
	}
	var existingRemote string
	if repository != nil {
		remote, err := repository.Remote("origin")
		if err != nil {
			logger.Fatal("Failed to get remote from repository: ", err)
		}
		existingRemote = remote.Config().URLs[0]
		logger.Info("Existing remote: ", existingRemote)
	}

	prompt := promptui.Prompt{
		Label:     "Repository",
		Default:   existingRemote,
		AllowEdit: true,
		Validate: func(s string) error {
			length := len(s)
			if length == 0 {
				return errors.New("repository cannot be empty")
			}

			if strings.HasPrefix(s, "http") {
				_, err := url.Parse(s)
				if err != nil {
					return errors.New("Invalid URL: " + err.Error())
				}
			}

			if !strings.HasSuffix(s, ".git") {
				return errors.New("repository URL should end with .git")
			}

			return nil
		},
	}
	result, err := prompt.Run()
	if err != nil {
		utils.IsInterrupt(&err)
		logger.Fatal("Prompt failed: ", err)
	}
	RepositoryUrl = result
}

func validateRepositoryUrl() {
	readUserInputRepositoryUrl()
	if RepositoryUrl == "" {
		logger.Fatal("Repository URL cannot be empty")
	}
	if strings.HasPrefix(RepositoryUrl, "http") {
		_, err := url.Parse(RepositoryUrl)
		if err != nil {
			logger.Fatal("Invalid URL: ", err)
		}
	}

	if !strings.HasSuffix(RepositoryUrl, ".git") {
		logger.Fatal("Repository URL should end with .git")
	}
	logger.Info("Using Repository: ", RepositoryUrl)
}
