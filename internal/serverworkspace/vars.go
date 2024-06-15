package serverworkspace

import (
	"github.com/arpanrec/netcli/internal/utils"
	"path"
)

var venvDir = path.Join(utils.GetHomeDir(), ".tmp", "serverworkspace-venv")
var venvEnvVars = make([]string, 0)
var LocalConfigPath = path.Join(utils.GetHomeDir(), ".tmp", "serverworkspace-local-config.json")
var basePythonPath string
var inventoryPath = path.Join(utils.GetHomeDir(), ".tmp", "serverworkspace-inventory.yaml")
