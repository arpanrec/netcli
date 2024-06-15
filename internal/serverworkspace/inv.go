package serverworkspace

import (
	"github.com/arpanrec/netcli/internal/logger"
	"github.com/arpanrec/netcli/internal/utils"
)

type inv struct {
	PythonPath string
}

func writeInventoryFile() {
	ansibleInv := inv{
		PythonPath: basePythonPath,
	}
	utils.WriteTextTemplateToFile("templates/serverworkspace-inventory.yaml.tmpl", "inventory", inventoryPath, ansibleInv)
	logger.Info("Inventory file written to: " + inventoryPath)
}
