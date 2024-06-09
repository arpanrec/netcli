package utils

import (
	"github.com/arpanrec/netcli/internal/logger"
	"os"
	"os/exec"
)

func BashExec(c *string) (string, error) {
	env := os.Environ()
	logger.Debug("Executing command: ", *c)
	cmd := exec.Command("/bin/bash", "-c", *c)
	cmd.Env = env
	out, err := cmd.CombinedOutput()
	return string(out), err
}
