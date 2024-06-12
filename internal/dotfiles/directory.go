package dotfiles

import (
	"os"
	"strings"

	"github.com/arpanrec/netcli/internal/logger"
	"github.com/arpanrec/netcli/internal/utils"
	gogit "github.com/go-git/go-git/v5"
	"github.com/manifoldco/promptui"
)

func readUserInputDirectory() {
	if gitDirectory != "" {
		return
	}
	prompt := promptui.Prompt{
		Label:   "Directory",
		Default: "~/.dotfiles",
		Validate: func(s string) error {
			return utils.ValidateDirectory(s, true, true)
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
	readUserInputDirectory()
	errValDir := utils.ValidateDirectory(gitDirectory, true, true)
	if errValDir != nil {
		logger.Fatal("Failed to validate directory: ", errValDir)
	}

	absPath, errAbsPath := utils.AbsPath(gitDirectory)
	if errAbsPath != nil {
		logger.Fatal("Failed to get absolute path: ", errAbsPath)
	}
	gitDirectory = absPath

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
	if !isSilent && !isCleanInstallProvided {
		_, errGitDirStat := os.Stat(gitDirectory)
		if errGitDirStat != nil {
			if os.IsNotExist(errGitDirStat) {
				return
			}
			logger.Fatal("Failed to get directory stat: ", errGitDirStat)
		}
		logger.Info("Do you want to clean install?")
		logger.Info("This will remove the git directory: ", gitDirectory, ", and clone the repository again.")
		options := []string{"No", "Yes"}
		prompt := promptui.Select{
			Label: "Clean Install?",
			Items: options,
			Searcher: func(input string, index int) bool {
				name := strings.Replace(strings.ToLower(options[index]), " ", "", -1)
				input = strings.Replace(strings.ToLower(input), " ", "", -1)
				return strings.Contains(name, input)
			},
		}
		_, result, err := prompt.Run()
		if err != nil {
			utils.IsInterrupt(&err)
			logger.Fatal("Prompt failed: ", err)
		}
		if result == "Yes" {
			isCleanInstall = true
		}
	}
	if isCleanInstall {
		logger.Info("Cleaning Git Directory: ", gitDirectory)
		err := os.RemoveAll(gitDirectory)
		if err != nil {
			logger.Fatal("Failed to clean Git Directory: ", err)
		}
	}
}
