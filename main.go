package main

import (
	"github.com/arpanrec/netcli/cmd"
	"github.com/arpanrec/netcli/internal/logger"
	"os"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		logger.Info("Failed to execute command", err)
		os.Exit(1)
	}
}
