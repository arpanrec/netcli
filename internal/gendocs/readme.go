package gendocs

import (
	"github.com/arpanrec/netcli/internal/constants"
	"github.com/arpanrec/netcli/internal/logger"
	"os"
	"path"
	"text/template"
)

type readme struct {
	DocsMdEp string
}

var readmeTemplate = `# Netcli

## [Usage]({{.DocsMdEp}})
`

func createReadme() {
	outputDirectoryBase := path.Base(outputDirectory)
	readmeLoc := path.Join(".", "README.md")

	tmpl, err := template.New("test").Parse(readmeTemplate)
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
	}

	err = tmpl.Execute(file, readmeMD)
}
