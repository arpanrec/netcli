package utils

import (
	"bytes"
	"os"
	"text/template"

	"github.com/arpanrec/netcli/assets"
	"github.com/arpanrec/netcli/internal/logger"
)

func getTextTemplate(templateFileName string, templateName string) *template.Template {
	a := &assets.Templates
	fileBytes, errFileBytes := a.ReadFile(templateFileName)
	if errFileBytes != nil {
		logger.Fatal("error reading template ", templateFileName, errFileBytes)
	}

	filesTmpl, errFilesTmpl := template.New(templateName).Parse(string(fileBytes))
	if errFilesTmpl != nil {
		logger.Fatal("error parsing template ", templateName, templateFileName, errFilesTmpl)
	}
	return filesTmpl
}
func GetTextTemplate(templateFileName string, templateName string, data any) string {

	filesTmpl := getTextTemplate(templateFileName, templateName)

	buf := &bytes.Buffer{}

	errExec := filesTmpl.Execute(buf, data)
	if errExec != nil {
		logger.Fatal("error executing template ", templateName, templateFileName, errExec)
	}
	return buf.String()
}

func WriteTextTemplate(templateFileName string, templateName string, dest string, data any) {
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