package serverworkspace

import (
	"os"

	"github.com/arpanrec/netcli/internal/logger"
)

func ifLocalConfigExist() {
	_, err := os.Stat(localConfigAbsPath)
	if err != nil {
		if os.IsNotExist(err) {
			writeErr := os.WriteFile(localConfigAbsPath, []byte("{}"), 0644)
			if writeErr != nil {
				logger.Fatal("Failed to write empty local config: ", writeErr)
			}
		}
	}
}
