package gendocs

import (
	"github.com/arpanrec/netcli/assets"
	"path"

	"github.com/arpanrec/netcli/internal/constants"
)

type readme struct {
	DocsMdEp string
	MainDesc string
	Version  string
}

func createReadme() {

	outputDirectoryBase := path.Base(OutputDirectory)
	readmeLoc := path.Join(".", "README.md")
	readmeMD := readme{
		DocsMdEp: path.Join(outputDirectoryBase, "netcli.md"),
		MainDesc: constants.NetCliShort + "\n\n" + constants.NetCliLong,
		Version:  constants.Version(),
	}
	assets.WriteTextTemplateToFile("templates/readme.md.tmpl", "readme", readmeLoc, readmeMD)
}
