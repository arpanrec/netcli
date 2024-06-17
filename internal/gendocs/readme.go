package gendocs

import (
	"path"

	"github.com/arpanrec/netcli/assets"
	"github.com/arpanrec/netcli/internal/vars"
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
		MainDesc: vars.NetCliShort + "\n\n" + vars.NetCliLong,
		Version:  vars.Version(),
	}
	assets.WriteTextTemplateToFile("templates/readme.md.tmpl", "readme", readmeLoc, readmeMD)
}
