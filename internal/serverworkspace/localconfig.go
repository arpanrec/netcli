package serverworkspace

import (
	"os"

	"github.com/arpanrec/netcli/internal/logger"
	"github.com/arpanrec/netcli/internal/utils"
)

func ifLocalConfigExist() {
	_, err := os.Stat(localConfigAbsPath)
	if err != nil {
		if os.IsNotExist(err) {
			createFileCmd := "echo {} | tee " + localConfigAbsPath
			ourCmd, errCmd := utils.BashExecEnv(&createFileCmd, &venvEnvVars)
			if errCmd != nil {
				logger.Fatal("Error creating local config file: ", errCmd, ourCmd)
			}
		}
	}
}
