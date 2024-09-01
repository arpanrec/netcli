package vars

import (
	"os"
	"sync"

	"github.com/arpanrec/netcli/internal/logger"
)

var homeDir string

var lockHomeDirFunc = &sync.Mutex{}

func GetHomeDir() string {

	if homeDir == "" {
		lockHomeDirFunc.Lock()
		defer lockHomeDirFunc.Unlock()
		if homeDir == "" {
			wd, wdErr := os.UserHomeDir()
			if wdErr != nil {
				logger.Fatal("Failed to get home gitDirectory: ", wdErr)
			}
			homeDir = wd
		}
	}
	return homeDir

}
