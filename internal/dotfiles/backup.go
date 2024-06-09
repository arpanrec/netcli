package dotfiles

import (
	"github.com/arpanrec/netcli/internal/logger"
	"strings"
)

func backup() {
	cmd := "ls-files"
	lsFiles := gitExec(&cmd)
	lsFilesArr := strings.Split(lsFiles, "\n")
	for _, file := range lsFilesArr {
		if file == "" {
			continue
		}
		logger.Info("Backing up file: ", file)
	}
}
