package utils

import (
	"github.com/arpanrec/netcli/internal/logger"
	"os"
)

func IsInterrupt(e *error) {
	if e == nil {
		return
	}
	if (*e).Error() == "^C" {
		logger.Fatal("IsInterrupt: Interrupted by user")
		os.Exit(1)
	}
	logger.Debug("IsInterrupt: Not an interrupt error: ", e)
}
