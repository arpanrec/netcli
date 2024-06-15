package constants

import (
	"github.com/arpanrec/netcli/internal/utils"
	"strings"
	"sync"
)

var lockVersionFunc = &sync.Mutex{}

var version *string

func Version() string {
	if version == nil {
		lockVersionFunc.Lock()
		defer lockVersionFunc.Unlock()
		if version == nil {
			versionFileContent := utils.GetTextFromTextTemplate("static/Version.txt", "version", nil)
			versionFileContentLines := strings.Split(versionFileContent, "\n")
			version = &versionFileContentLines[0]
		}
	}
	return *version
}
