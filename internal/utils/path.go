package utils

import (
	"errors"
	"path/filepath"
	"strings"
)

func AbsPath(p string) (string, error) {

	if strings.HasPrefix(p, "~/") || p == "~" {
		if strings.HasPrefix(p, "~/") {
			p = strings.Replace(p, "~/", homeDir+"/", 1)
		}
		if p == "~" {
			p = homeDir
		}
	}
	if strings.Contains(p, "$") {
		cmd := "echo " + p
		out, err := BashExec(&cmd)
		if err != nil {
			return "", errors.New("failed to get absolute path using shell echo: , " + err.Error())
		}
		p = strings.TrimSpace(out)
	}
	absPath, errAbs := filepath.Abs(p)
	if errAbs != nil {
		return "", errors.New("failed to get absolute path of the file, " + errAbs.Error())
	}
	p = absPath
	return p, nil
}
