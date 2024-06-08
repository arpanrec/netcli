package dotfiles

import (
	"errors"
	"github.com/arpanrec/netcli/internal/logger"
	"github.com/arpanrec/netcli/internal/utils"
	"github.com/go-git/go-git/v5"
	"github.com/manifoldco/promptui"
	"os"
	"path/filepath"
)

func readUserInputDirectory() {
	if directory != "" {
		return
	}
	prompt := promptui.Prompt{
		Label: "Directory",
		Validate: func(s string) error {
			length := len(s)
			if length == 0 {
				return errors.New("directory cannot be empty")
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
		utils.IsInterrupt(err)
		logger.Fatal("Prompt failed: ", err)
	}
	directory = result
}

func validateDirectoryAndLoadRepo() {
	utils.ExpectingCleanPath(&directory)
	errAbsPath := utils.AbsPath(&directory)
	if errAbsPath != nil {
		logger.Fatal("Failed to get absolute path: ", errAbsPath)
	}
	stat, err := os.Stat(directory)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		logger.Fatal("Directory does not exist: ", directory)
	}
	logger.Debug("Directory info: ", stat)
	if !stat.IsDir() {
		logger.Fatal("Path is not a directory: ", directory)
	}

	r, errR := git.PlainOpen(directory)
	if errR != nil {
		logger.Fatal("Directory :", directory, ", is not a git repository. Error: ", errR)
	}
	repository = r
	logger.Info("Repository loaded from directory: ", directory)
}
