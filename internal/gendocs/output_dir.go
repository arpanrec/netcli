package gendocs

import (
	"os"
	"path/filepath"

	"github.com/arpanrec/netcli/internal/logger"
	"github.com/arpanrec/netcli/internal/utils"
)

func outputDir() {
	if OutputDirectory == "" {
		logger.Fatal("output directory cannot be empty")
	}

	absPath, errAbsPath := filepath.Abs(OutputDirectory)
	if errAbsPath != nil {
		logger.Fatal("error getting absolute path of output directory")
	}
	OutputDirectory = absPath

	errValidate := utils.ValidateDirectory(OutputDirectory, true, true)
	if errValidate != nil {
		logger.Fatal("error validating output directory: " + errValidate.Error())
	}

	errRemove := os.RemoveAll(OutputDirectory)
	if errRemove != nil {
		logger.Fatal("error removing existing files in output directory: " + errRemove.Error())
	}
	logger.Info("Creating output directory: ", OutputDirectory)
	errMkdir := os.MkdirAll(OutputDirectory, 0755)
	if errMkdir != nil {
		logger.Fatal("error creating output directory: " + errMkdir.Error())
	}
}
