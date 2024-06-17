package serverworkspace

import "github.com/arpanrec/netcli/internal/logger"

func preCheck() {
	if rawArgs == "" && rawArgsProvided {
		logger.Fatal("Raw arguments provided but not found")
	}

}
