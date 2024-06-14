package main

import (
	"github.com/arpanrec/netcli/cmd"
	"github.com/arpanrec/netcli/internal/logger"
)

func main() {
	logger.SetUpLogger()
	cmd.Execute()
	logger.Sync()
}
