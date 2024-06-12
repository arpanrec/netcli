package gendocs

import (
	"os"
	"path"
	"text/template"

	"github.com/arpanrec/netcli/internal/constants"
	"github.com/arpanrec/netcli/internal/logger"
)

type readme struct {
	DocsMdEp string
	MainDesc string
}

var readmeTemplate = `# Netcli

{{.MainDesc}}

## [Usage]({{.DocsMdEp}})
`

func createReadme() {
	outputDirectoryBase := path.Base(outputDirectory)
	readmeLoc := path.Join(".", "README.md")

	tmpl, err := template.New("README").Parse(readmeTemplate)
	if err != nil {
		logger.Panic("error parsing template", err)
	}
	file, errCreate := os.Create(readmeLoc)
	if errCreate != nil {
		logger.Panic("error creating file", errCreate)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logger.Panic("error closing file", err)
		}
	}(file)

	readmeMD := readme{
		DocsMdEp: path.Join(outputDirectoryBase, constants.NetCliUse+".md"),
		MainDesc: constants.NetCliShort + "\n\n" + constants.NetCliLong,
	}

	errExec := tmpl.Execute(file, readmeMD)
	if errExec != nil {
		logger.Panic("error executing template", errExec)
	}
}
