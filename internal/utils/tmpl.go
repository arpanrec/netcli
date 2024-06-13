package utils

import (
	"bytes"
	"embed"
	"os"
	"strings"
	"text/template"

	"github.com/arpanrec/netcli/assets"
	"github.com/arpanrec/netcli/internal/logger"
)

func getTextTemplate(templateFileName string, templateName string) *template.Template {
	var a *embed.FS
	if strings.HasPrefix(templateFileName, "templates/") {
		a = &assets.Templates
	} else if strings.HasPrefix(templateFileName, "static/") {
		a = &assets.StaticFiles
	}
	if a == nil {
		logger.Fatal("template file name should start with templates/ or static/")
		os.Exit(1)
	}

	fileBytes, errFileBytes := a.ReadFile(templateFileName)
	if errFileBytes != nil {
		logger.Fatal("error reading template\n", templateFileName, "\n", errFileBytes)
	}

	filesTmpl, errFilesTmpl := template.New(templateName).Parse(string(fileBytes))
	if errFilesTmpl != nil {
		logger.Fatal("error parsing template ", templateName, templateFileName, errFilesTmpl)
	}
	return filesTmpl
}
func GetTextFromTextTemplate(templateFileName string, templateName string, data any) string {

	filesTmpl := getTextTemplate(templateFileName, templateName)

	buf := &bytes.Buffer{}

	errExec := filesTmpl.Execute(buf, data)
	if errExec != nil {
		logger.Fatal("error executing template ", templateName, templateFileName, errExec)
	}
	return buf.String()
}

func WriteTextTemplateToFile(templateFileName string, templateName string, dest string, data any) {
	filesTmpl := getTextTemplate(templateFileName, templateName)
	file, errCreate := os.Create(dest)
	if errCreate != nil {
		logger.Fatal("error creating file ", errCreate)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logger.Fatal("error closing file", err)
		}
	}(file)

	errExec := filesTmpl.Execute(file, data)
	if errExec != nil {
		logger.Fatal("error executing template", templateName, templateFileName, errExec)
	}
}
