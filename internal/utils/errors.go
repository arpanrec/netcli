package utils

import (
	"github.com/arpanrec/netcli/internal/logger"
)

func IsInterrupt(e *error) {
	if e == nil {
		return
	}
	if (*e).Error() == "^C" {
		logger.Fatal("IsInterrupt: Interrupted by user ", *e)
	}
	logger.Debug("IsInterrupt: Not an interrupt error: ", e)
}
