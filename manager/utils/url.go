package utils

import (
	"errors"
	"strings"
)

func CheckPathService(module string, path string) error {
	if !strings.HasPrefix(path, "/") {
		return errors.New("path no start with prefix \"/\"")
	}

	checks := strings.Split(path, "/")
	if checks[1] != module {
		return errors.New("path no has module code")
	}

	return nil
}
