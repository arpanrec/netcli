package serverworkspace

import (
	"github.com/arpanrec/netcli/internal/logger"
	"github.com/arpanrec/netcli/internal/utils"
)

func run() {
	var cmd = "ansible-playbook arpanrec.nebula.server_workspace" +
		" --inventory " + inventoryPath +
		" --extra-vars @" + localConfigAbsPath
	if rawArgs != "" {
		cmd = cmd + " " + rawArgs
	} else {
		cmd = cmd + " --tags "
		if nodeJs {
			cmd = cmd + "nodejs,"
		}
		if golang {
			cmd = cmd + "go,"
		}
		if java {
			cmd = cmd + "java,"
		}
		if terminal {
			cmd = cmd + "terminal,"
		}
		if terraform {
			cmd = cmd + "terraform,"
		}
		if vault {
			cmd = cmd + "vault,"
		}
		if pulumi {
			cmd = cmd + "pulumi,"
		}
		if bws {
			cmd = cmd + "bws,"
		}
		cmd = cmd[:len(cmd)-1]
	}
	logger.Info("Running command: " + cmd)
	logger.Info("Please wait, this may take a while...")
	out, err := utils.BashExecEnv(&cmd, &venvEnvVars)
	if err != nil {
		logger.Fatal("Failed to run command: ", out, err)
	}
	logger.Info(out)
}
