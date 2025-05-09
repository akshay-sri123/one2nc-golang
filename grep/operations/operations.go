package operations

import (
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
