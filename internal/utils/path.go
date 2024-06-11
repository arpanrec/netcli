package utils

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func AbsPath(p string) (string, error) {

	if strings.HasPrefix(p, "~/") || p == "~" ||
		strings.HasSuffix(p, "/~") || strings.Contains(p, "/~/") {
		homeDir, errHomeDir := os.UserHomeDir()
		if errHomeDir != nil {
			return "", errors.New("failed to get user home directory, " + errHomeDir.Error())
		}
		if strings.HasPrefix(p, "~/") {
			p = strings.Replace(p, "~/", homeDir+"/", 1)
		}
		if p == "~" {
			p = homeDir
		}
		if strings.HasSuffix(p, "/~") {
			p = strings.Replace(p, "/~", "/"+homeDir, 1)
		}

		if strings.Contains(p, "/~/") {
			p = strings.ReplaceAll(p, "/~/", "/"+homeDir+"/")

		}
	}
	if strings.Contains(p, "$") {
		var envVars []string
		copy(envVars, os.Environ())
		cmd := exec.Command("/bin/bash", "-c", "realpath "+p)
		cmd.Env = envVars
		out, err := cmd.CombinedOutput()
		if err != nil {
			return "", errors.New("failed to get absolute path of SSH key using realpath, " + err.Error())
		}
		p = strings.TrimSpace(string(out))
	}
	absPath, errAbs := filepath.Abs(p)
	if errAbs != nil {
		return "", errors.New("failed to get absolute path of SSH key, " + errAbs.Error())
	}
	p = absPath
	return p, nil
}
