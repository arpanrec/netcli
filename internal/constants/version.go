package constants

import (
	"github.com/arpanrec/netcli/internal/utils"
	"strings"
)

func Version() string {
	versionFileContent := utils.GetTextFromTextTemplate("static/Version.txt", "version", nil)
	versionFileContentLines := strings.Split(versionFileContent, "\n")
	return versionFileContentLines[0]
}
