package dotfiles

import (
	"errors"
	"github.com/arpanrec/netcli/internal/logger"
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
			return nil
		},
	}
	result, err := prompt.Run()
	if err != nil {
		logger.Fatal("Prompt failed: ", err)
	}
	if result == "" {
		logger.Fatal("Directory cannot be empty")
	}
	directory = result

}

func validateDirectoryAndLoadRepo() {
	if !filepath.IsAbs(directory) {
		absPath, errAbsPath := filepath.Abs(directory)
		if errAbsPath != nil {
			logger.Fatal("Failed to get absolute path: ", errAbsPath)
		}
		directory = absPath
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
		logger.Fatal("Directory is not a git repository: ", errR)
	}
	repository = r
	logger.Info("Repository loaded from directory: ", directory)
}
