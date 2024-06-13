package gendocs

import (
	"path"

	"github.com/arpanrec/netcli/internal/constants"
	"github.com/arpanrec/netcli/internal/utils"
)

type readme struct {
	DocsMdEp string
	MainDesc string
}

func createReadme() {

	outputDirectoryBase := path.Base(OutputDirectory)
	readmeLoc := path.Join(".", "README.md")
	readmeMD := readme{
		DocsMdEp: path.Join(outputDirectoryBase, "netcli.md"),
		MainDesc: constants.NetCliShort + "\n\n" + constants.NetCliLong,
	}
	utils.WriteTextTemplate("templates/readme.md.tmpl", "readme", readmeLoc, readmeMD)
}
