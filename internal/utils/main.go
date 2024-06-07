package utils

import (
	"github.com/arpanrec/netcli/internal/logger"
	"golang.org/x/term"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ReadChars(num int) string {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		logger.Fatal("Failed to switch to raw mode for stdin", err)
	}
	defer func(fd int, oldState *term.State) {
		err := term.Restore(fd, oldState)
		if err != nil {
			logger.Fatal("Failed to restore terminal to previous state", err)
		}
	}(int(os.Stdin.Fd()), oldState)
	b := make([]byte, num)
	_, err = os.Stdin.Read(b)
	if err != nil {
		logger.Fatal("Failed to read from stdin", err)
	}
	return string(b)
}

func AbsPath(p *string) error {

	if strings.HasPrefix(*p, "~/") || *p == "~" ||
		strings.HasSuffix(*p, "/~") || strings.Contains(*p, "/~/") {
		homeDir, errHomeDir := os.UserHomeDir()
		if errHomeDir != nil {
			return errHomeDir
		}
		if strings.HasPrefix(*p, "~/") {
			*p = strings.Replace(*p, "~/", homeDir+"/", 1)
		}
		if *p == "~" {
			*p = homeDir
		}
		if strings.HasSuffix(*p, "/~") {
			*p = strings.Replace(*p, "/~", "/"+homeDir, 1)
		}

		if strings.Contains(*p, "/~/") {
			*p = strings.ReplaceAll(*p, "/~/", "/"+homeDir+"/")

		}
	}
	if strings.Contains(*p, "$") {
		envVars := os.Environ()
		cmd := exec.Command("/bin/bash", "-c", "realpath "+*p)
		cmd.Env = envVars
		out, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		*p = strings.TrimSpace(string(out))
	}
	absPath, errAbs := filepath.Abs(*p)
	if errAbs != nil {
		return errAbs
	}
	*p = absPath
	return nil

}
