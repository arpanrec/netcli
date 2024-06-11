package utils

import (
	"os"
	"os/exec"
)

func BashExec(c *string) (string, error) {
	env := os.Environ()
	cmd := exec.Command("/bin/bash", "-c", *c)
	cmd.Env = env
	out, err := cmd.CombinedOutput()
	return string(out), err
}
