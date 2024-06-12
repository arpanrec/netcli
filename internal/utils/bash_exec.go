package utils

import (
	"os"
	"os/exec"
)

func BashExec(c *string) (string, error) {
	env := os.Environ()
	return BashExecEnv(c, &env)
}

func BashExecEnv(c *string, env *[]string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", *c)
	cmd.Env = *env
	out, err := cmd.CombinedOutput()
	return string(out), err
}
