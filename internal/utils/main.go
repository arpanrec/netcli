package utils

import (
	"github.com/arpanrec/netcli/internal/logger"
	"golang.org/x/term"
	"os"
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
