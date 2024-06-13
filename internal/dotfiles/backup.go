package dotfiles

import (
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/arpanrec/netcli/internal/logger"
	"github.com/arpanrec/netcli/internal/utils"
	"github.com/manifoldco/promptui"
)

func backup() {
	logger.Info("Backup command called")
	lsFilesArr := tempCheckOutLsFiles()
	readUserInputBackupDirectory()
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

		backupFile := path.Join(BackupDir, file)
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

func tempCheckOutLsFiles() []string {
	dirPath, err := os.MkdirTemp("", "netcli-backup-*")
	if err != nil {
		logger.Fatal("Failed to create temporary directory: ", err)
	}
	logger.Info("Created temporary directory: ", dirPath)
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			logger.Warn("Failed to remove temporary directory: ", err)
		} else {
			logger.Info("Removed temporary directory: ", path)
		}
	}(dirPath)

	cmd := "checkout"
	_ = gitExecWd(&cmd, &dirPath)
	cmd = "ls-files"
	lsFiles := gitExecWd(&cmd, &dirPath)
	return strings.Split(lsFiles, "\n")
}

func readUserInputBackupDirectory() {
	if BackupDir != "" {
		return
	}

	defaultBackupDir := path.Join(backupDirRoot, strconv.Itoa(int(time.Now().Unix())))
	if isSilent {
		BackupDir = defaultBackupDir
		logger.Info("Using default backup directory: ", BackupDir)
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
	BackupDir = result
	logger.Info("Backup directory set to: ", BackupDir)
}
