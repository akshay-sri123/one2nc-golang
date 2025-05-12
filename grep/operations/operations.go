package operations

import (
	"errors"
	"os"
	"strings"
)

func searchSubstring(text string, substring string) []string {
	var resultSet []string
	for _, line := range strings.Split(text, "\n") {
		if strings.Contains(line, substring) {
			resultSet = append(resultSet, line)
		}
	}
	return resultSet
}

func readFromFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func checkIfFileExists(filename string) bool {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func checkIfFileOrDir(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}
	if info.IsDir() {
		return false
	}
	return true
}

func checkFilePermissions(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}
	if info.Mode().Perm()&0444 != 0444 {
		return false
	}
	return true
}
