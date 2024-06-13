package gendocs

import (
	"path"

	"github.com/arpanrec/netcli/internal/constants"
	"github.com/arpanrec/netcli/internal/utils"
)

type webrun struct {
	Version string
}

func createWebRunScript() {
	readmeLoc := path.Join(".", "web_run.sh")
	webrunSh := webrun{
		Version: constants.Version(),
	}
	utils.WriteTextTemplate("templates/web_run.sh.impl", "web_run", readmeLoc, webrunSh)
}
