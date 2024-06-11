package gendocs

import (
	"github.com/arpanrec/netcli/internal/logger"
)

func preChecks() {

	if outputDirectoryProvided && outputDirectory == "" {
		logger.Fatal("output directory cannot be empty")
	}

}
