package serverworkspace

import (
	"github.com/arpanrec/netcli/internal/logger"
	"github.com/arpanrec/netcli/internal/utils"
)

func run() {
	var cmd = "ansible-playbook arpanrec.nebula.server_workspace" +
		" --inventory " + inventoryPath +
		" --extra-vars @" + LocalConfigPath
	if RawArgs != "" {
		cmd = cmd + " " + RawArgs
	} else {
		cmd = cmd + " --tags "
		if NodeJs {
			cmd = cmd + "nodejs,"
		}
		if Go {
			cmd = cmd + "go,"
		}
		if Java {
			cmd = cmd + "java,"
		}
		if Terminal {
			cmd = cmd + "terminal,"
		}
		if Terraform {
			cmd = cmd + "terraform,"
		}
		if Vault {
			cmd = cmd + "vault,"
		}
		if Pulumi {
			cmd = cmd + "pulumi,"
		}
		if BWS {
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
