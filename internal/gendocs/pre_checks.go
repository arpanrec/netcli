package gendocs

import "github.com/arpanrec/netcli/internal/logger"

func preChecks() {

	if outputDirectoryProvided && OutputDirectory == "" {
		logger.Fatal("output directory cannot be empty")
	}

}
