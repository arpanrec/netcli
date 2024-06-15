package serverworkspace

import "github.com/arpanrec/netcli/internal/logger"

func preCheck() {
	if RawArgs == "" && RawArgsProvided {
		logger.Fatal("Raw arguments provided but not found")
	}

}
