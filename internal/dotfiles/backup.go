package dotfiles

import (
	"github.com/arpanrec/netcli/internal/logger"
	"github.com/arpanrec/netcli/internal/utils"
	"github.com/manifoldco/promptui"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func backup() {
	readUserInputBackupDirectory()
	cmd := "ls-files"
	lsFiles := gitExec(&cmd)
	lsFilesArr := strings.Split(lsFiles, "\n")
	for _, file := range lsFilesArr {
		if file == "" {
			continue
		}
		sourceFile := path.Join(workTreeDir, file)
		fileStat, fileStatErr := os.Stat(sourceFile)
		if fileStatErr != nil {
			if os.IsNotExist(fileStatErr) {
				logger.Info("File does not exist: ", sourceFile)
				continue
			}
			logger.Fatal("Error getting file stat: ", fileStatErr, sourceFile)
		}

		if !fileStat.Mode().IsRegular() {
			logger.Warn("How the hell a non-regular file is in the git repo?", sourceFile)
			continue
		}

		if fileStat.IsDir() {
			logger.Fatal("How the hell a directory is in the git repo?", sourceFile)
		}

		backupFile := path.Join(backupDir, file)
		backupFileDir := path.Dir(backupFile)
		if _, err := os.Stat(backupFileDir); os.IsNotExist(err) {
			logger.Info("Creating directory: ", backupFileDir)
			if err := os.MkdirAll(backupFileDir, 0755); err != nil {
				logger.Fatal("Failed to create directory: ", backupFileDir, err)
			}
		}
		logger.Info("Copying file: ", sourceFile, " to: ", backupFile)
		copyCMD := "cp '" + sourceFile + "' '" + backupFile + "'"
		_, errBash := utils.BashExec(&copyCMD)
		if errBash != nil {
			logger.Fatal("Failed to copy file: ", sourceFile, " to: ", backupFile, errBash)
		}

	}
}

func readUserInputBackupDirectory() {
	if backupDir != "" {
		return
	}

	defaultBackupDir := path.Join(backupDirRoot, strconv.Itoa(int(time.Now().Unix())))
	if isSilent {
		backupDir = defaultBackupDir
		logger.Info("Using default backup directory: ", backupDir)
		return
	}
	prompt := promptui.Prompt{
		Label:     "Backup Directory",
		Default:   defaultBackupDir,
		AllowEdit: true,
		Validate: func(s string) error {
			return utils.ValidateDirectory(s, true, false)
		},
	}
	result, err := prompt.Run()
	if err != nil {
		utils.IsInterrupt(&err)
		logger.Fatal("Prompt failed: ", err)
	}
	backupDir = result
	logger.Info("Backup directory set to: ", backupDir)
}
