package dotfiles

import (
	"fmt"
	"github.com/arpanrec/netcli/internal/logger"
	"github.com/arpanrec/netcli/internal/utils"
)

func gitExec(command *string) string {
	commandFormat := fmt.Sprintf("git --git-dir=%s --work-tree=%s %s", gitDirectory, workTreeDir, *command)
	out, cmdErr := utils.BashExec(&commandFormat)
	if cmdErr != nil {
		logger.Fatal("Failed to execute command: ", out, cmdErr)
	}
	return out
}