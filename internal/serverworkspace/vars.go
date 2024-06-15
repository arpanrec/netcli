package serverworkspace

import (
	"github.com/arpanrec/netcli/internal/utils"
	"path"
)

var venvDir = path.Join(utils.GetHomeDir(), ".tmp", "serverworkspace-venv")
var venvEnvVars = make([]string, 0)
