package utils

import (
	"errors"
	"os"
	"path"
	"strings"
)

func ValidateDirectory(p string, okIsNotExist bool, okIfFilesExists bool) error {
	if p == "" {
		return errors.New("directory path is empty")
	}
	if strings.HasSuffix(p, "/") {
		return errors.New("directory should not end with a slash")
	}
	cleanPath := path.Clean(p)
	if cleanPath != p {
		return errors.New("invalid path: " + p +
			". Clean path will look like: " + cleanPath +
			", path is not clean. Check https://pkg.go.dev/path#Clean for more details")
	}

	absPath, absPathErr := AbsPath(p)
	if absPathErr != nil {
		return errors.New("failed to get absolute path: " + absPathErr.Error())
	}
	p = absPath

	stat, statErr := os.Stat(p)
	if statErr != nil {
		if os.IsNotExist(statErr) && okIsNotExist {
			return nil
		}
		return errors.New("failed to get stat: " + statErr.Error())
	}
	if !stat.IsDir() {
		return errors.New("path is not a directory: " + p)
	}

	lsCmd := "ls -aA " + p
	lsOutput, bashErr := BashExec(&lsCmd)
	if bashErr != nil {
		return errors.New("failed to list files in directory: " + lsOutput + "\n" + bashErr.Error())
	}
	if lsOutput != "" && !okIfFilesExists {
		return errors.New("directory is not empty: " + p)
	}

	return nil
}

func ValidateFile(p string, okIsNotExist bool) error {
	if p == "" {
		return errors.New("file path is empty")
	}
	if strings.HasSuffix(p, "/") {
		return errors.New("file path should not end with a slash")
	}
	cleanPath := path.Clean(p)
	if cleanPath != p {
		return errors.New("invalid path: " + p +
			". Clean path will look like: " + cleanPath +
			", path is not clean. Check https://pkg.go.dev/path#Clean for more details")
	}

	absPath, absPathErr := AbsPath(p)
	if absPathErr != nil {
		return errors.New("failed to get absolute path: " + absPathErr.Error())
	}
	p = absPath

	stat, statErr := os.Stat(p)
	if statErr != nil {
		if os.IsNotExist(statErr) && okIsNotExist {
			return nil
		}
		return errors.New("failed to get stat: " + statErr.Error())
	}
	if stat.IsDir() {
		return errors.New("path is a directory: " + p)
	}

	return nil
}
