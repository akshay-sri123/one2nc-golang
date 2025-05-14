package operations

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type FlagOperations struct {
	FilterString   string
	FilesToProcess string
	StdInput       string
}

type GrepResult struct {
	filename        string
	matchingResults []string
}

func (grepResult GrepResult) printResults() {
	for _, result := range grepResult.matchingResults {
		fmt.Println(result)
	}
}

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

func checkFile(filename string, command string) error {
	errOutput := ""
	var err error
	if !checkIfFileExists(filename) {
		errOutput = fmt.Sprintf("%s: %s: read: No such file or directory\n", command, filename)
		err = errors.New(errOutput)
	}

	if !checkIfFileOrDir(filename) {
		errOutput = fmt.Sprintf("%s: %s: open: Is a directory\n", command, filename)
		err = errors.New(errOutput)
	}

	if !checkFilePermissions(filename) {
		errOutput = fmt.Sprintf("%s: %s: open: Permission denied\n", command, filename)
		err = errors.New(errOutput)
	}
	return err
}

func RunOperation(flagOperations FlagOperations, command string) {
	var grepResult GrepResult
	if flagOperations.FilesToProcess != "" {
		err := checkFile(flagOperations.FilesToProcess, command)
		if err != nil {
			fmt.Println(err.Error())
		}
		fileText, err := readFromFile(flagOperations.FilesToProcess)
		if err != nil {
			fmt.Println(err.Error())
		}
		grepResult.filename = flagOperations.FilesToProcess

		grepResult.matchingResults = searchSubstring(fileText, flagOperations.FilterString)
	} else if flagOperations.StdInput != "" {
		grepResult.matchingResults = searchSubstring(flagOperations.StdInput, flagOperations.FilterString)
	}
	grepResult.printResults()
}
