package util

import (
	"fmt"
	"os"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)

	if err == nil {
		return true
	}

	return false
}


func HomeFile(file string) (string, error) {
	homedir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", homedir, file), nil
}