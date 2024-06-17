package serverworkspace

import (
	"path"

	"github.com/arpanrec/netcli/internal/vars"
)

var venvDir = path.Join(vars.GetHomeDir(), ".tmp", "serverworkspace-venv")
var venvEnvVars = make([]string, 0)
var LocalConfigPath = path.Join(vars.GetHomeDir(), ".tmp", "serverworkspace-local-config.json")
var basePythonPath string
var inventoryPath = path.Join(vars.GetHomeDir(), ".tmp", "serverworkspace-inventory.yaml")
