package dotfiles

import (
	"errors"
	"github.com/arpanrec/netcli/internal/logger"
	"github.com/arpanrec/netcli/internal/utils"
	"github.com/manifoldco/promptui"
	"net/url"
	"strings"
)

func readUserInputRepositoryUrl() {
	if repositoryUrl != "" {
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
	repositoryUrl = result
}

func validateRepositoryUrl() {
	if repositoryUrl == "" {
		logger.Fatal("Repository URL cannot be empty")
	}
	if strings.HasPrefix(repositoryUrl, "http") {
		_, err := url.Parse(repositoryUrl)
		if err != nil {
			logger.Fatal("Invalid URL: ", err)
		}
	}

	if !strings.HasSuffix(repositoryUrl, ".git") {
		logger.Fatal("Repository URL should end with .git")
	}
	logger.Info("Using Repository: ", repositoryUrl)
}
