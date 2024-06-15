package serverworkspace

import (
	"github.com/arpanrec/netcli/assets"
	"github.com/arpanrec/netcli/internal/logger"
)

type inv struct {
	PythonPath string
}

func writeInventoryFile() {
	ansibleInv := inv{
		PythonPath: basePythonPath,
	}
	assets.WriteTextTemplateToFile("templates/serverworkspace-inventory.yaml.tmpl", "inventory", inventoryPath, ansibleInv)
	logger.Info("Inventory file written to: " + inventoryPath)
}
