package dotfiles

import (
	"errors"
	"github.com/arpanrec/netcli/internal/logger"
	"github.com/arpanrec/netcli/internal/utils"
	gogit "github.com/go-git/go-git/v5"
	"github.com/manifoldco/promptui"
	"os"
	"path/filepath"
)

func readUserInputDirectory() {
	if gitDirectory != "" {
		return
	}
	prompt := promptui.Prompt{
		Label:   "Directory",
		Default: "~/.dotfiles",
		Validate: func(s string) error {
			length := len(s)
			if length == 0 {
				return errors.New("gitDirectory cannot be empty")
			}
			cleanPath := filepath.Clean(s)
			if cleanPath != s {
				return errors.New("Invalid path: " + s +
					". Clean path will look like: " + cleanPath +
					", path is not clean. Check https://pkg.go.dev/path#Clean for more details")
			}
			return nil
		},
	}
	result, err := prompt.Run()
	if err != nil {
		utils.IsInterrupt(&err)
		logger.Fatal("Prompt failed: ", err)
	}
	gitDirectory = result
}

func validateDirectoryAndLoadRepo() {
	utils.ExpectingCleanPath(&gitDirectory)
	errAbsPath := utils.AbsPath(&gitDirectory)
	if errAbsPath != nil {
		logger.Fatal("Failed to get absolute path: ", errAbsPath)
	}
	logger.Info("Directory Absolute path: ", gitDirectory)
	cleanInstall()
	stat, err := os.Stat(gitDirectory)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		logger.Fatal("Directory does not exist: ", gitDirectory)
	}
	logger.Debug("Directory info: ", stat)
	if !stat.IsDir() {
		logger.Fatal("Path is not a directory: ", gitDirectory)
	}

	r, errR := gogit.PlainOpenWithOptions(gitDirectory, &gogit.PlainOpenOptions{})
	if errR != nil {
		logger.Fatal("Directory :", gitDirectory, ", is not a git repository. Error: ", errR)
	}
	repository = r
	logger.Info("Repository loaded from gitDirectory: ", gitDirectory)
}

func cleanInstall() {
	if isCleanInstall {
		logger.Info("Cleaning Git Directory: ", gitDirectory)
		err := os.RemoveAll(gitDirectory)
		if err != nil {
			logger.Fatal("Failed to clean Git Directory: ", err)
		}
	}
}
