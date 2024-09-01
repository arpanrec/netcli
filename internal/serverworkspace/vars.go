package serverworkspace

import (
	"path"

	"github.com/arpanrec/netcli/internal/vars"
)

var venvDir = path.Join(vars.GetHomeDir(), ".tmp", "serverworkspace-venv")
var venvEnvVars = make([]string, 0)
var localConfigPath = path.Join(".tmp", "serverworkspace-local-config.json")
var localConfigAbsPath = path.Join(vars.GetHomeDir(), localConfigPath)
var basePythonPath string
var inventoryPath = path.Join(vars.GetHomeDir(), ".tmp", "serverworkspace-inventory.yaml")
