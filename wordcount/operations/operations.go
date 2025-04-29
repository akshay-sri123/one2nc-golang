package operations

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type FlagOperations struct {
	CLines bool
	CWords bool
	CChars bool
}

type OperationResults struct {
	NLines   int
	NWords   int
	NChars   int
	Filename string
}

func countLines(text string) int {
	splitLines := strings.Split(text, "\n")
	return len(splitLines) - 1
}

func countWords(text string) int {
	splitWords := strings.Fields(text)
	return len(splitWords)
}

func countCharacters(text string) int {
	charCount := 0
	splitWords := strings.Split(text, " ")

	for _, word := range splitWords {
		charCount = charCount + len(word)
	}
	return (charCount + len(splitWords) - 1)
}

func checkIfFileExists(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
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

func readTextFromFile(filename string) string {
	data, err := os.ReadFile(filename)
	check(err)

	return string(data)
}

func CountOperations(executeOperations FlagOperations, filesToProcess []string, command string) []OperationResults {
	var operationResultsList []OperationResults
	for _, file := range filesToProcess {
		var operationResults OperationResults
		err := checkFile(file, command)
		if err != nil {
			fmt.Print(err.Error())
			continue
		}
		text := readTextFromFile(file)

		if executeOperations.CLines {
			operationResults.NLines = countLines(text)
		}

		if executeOperations.CWords {
			operationResults.NWords = countWords(text)
		}

		if executeOperations.CChars {
			operationResults.NChars = countCharacters(text)
		}

		if !executeOperations.CLines && !executeOperations.CWords && !executeOperations.CChars {
			operationResults.NLines = countLines(text)
			operationResults.NWords = countWords(text)
			operationResults.NChars = countCharacters(text)
		}

		operationResults.Filename = file
		operationResultsList = append(operationResultsList, operationResults)
	}
	return operationResultsList
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GenerateOutput(operationResults []OperationResults, executeOperations FlagOperations) string {
	var finalResult string
	for _, opRes := range operationResults {
		var output string
		if executeOperations.CLines {
			output += fmt.Sprintf("%8d", opRes.NLines)
		}

		if executeOperations.CWords {
			output += fmt.Sprintf("%8d", opRes.NWords)
		}

		if executeOperations.CChars {
			output += fmt.Sprintf("%8d", opRes.NChars)
		}

		if !executeOperations.CLines && !executeOperations.CWords && !executeOperations.CChars {
			output += fmt.Sprintf("%8d", opRes.NLines)
			output += fmt.Sprintf("%8d", opRes.NWords)
			output += fmt.Sprintf("%8d", opRes.NChars)
		}
		output += fmt.Sprint(" " + opRes.Filename)
		finalResult += output + "\n"
	}
	return finalResult
}
