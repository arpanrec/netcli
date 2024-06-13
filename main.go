package main

import (
	"github.com/arpanrec/netcli/cmd"
	"github.com/arpanrec/netcli/internal/logger"
)

func main() {
	cmd.Execute()
}

func init() {
	logger.SetUpLogger()
}
