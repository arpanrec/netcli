package gendocs

import (
	"github.com/arpanrec/netcli/internal/logger"
	"os"
	"text/template"
)

type readme struct {
	Material string
	Count    uint
}

var readmeTemplate = `# Netcli

This is a CLI application for managing network devices.
`

func createReadme() {
	readmeMD := readme{"wool", 17}
	tmpl, err := template.New("test").Parse(readmeTemplate)
	if err != nil {
		panic(err)
	}
	file, _ := os.Create("README.md")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logger.Panic("error closing file")
		}
	}(file)

	err = tmpl.Execute(file, readmeMD)
}
