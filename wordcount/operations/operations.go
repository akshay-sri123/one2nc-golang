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

func readTextFromFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func countOperation(executeOperations FlagOperations, file string, command string) (OperationResults, error) {
	var operationResult OperationResults
	err := checkFile(file, command)
	if err != nil {
		return operationResult, err
	}
	text, err := readTextFromFile(file)
	if err != nil {
		return operationResult, err
	}

	if executeOperations.CLines {
		operationResult.NLines = countLines(text)
	}

	if executeOperations.CWords {
		operationResult.NWords = countWords(text)
	}

	if executeOperations.CChars {
		operationResult.NChars = countCharacters(text)
	}

	if !executeOperations.CLines && !executeOperations.CWords && !executeOperations.CChars {
		operationResult.NLines = countLines(text)
		operationResult.NWords = countWords(text)
		operationResult.NChars = countCharacters(text)
	}

	operationResult.Filename = file
	return operationResult, nil
}

func generateOutput(operationResults []OperationResults, executeOperations FlagOperations) {
	var finalResult string
	for _, opRes := range operationResults {
		if opRes.NLines == 0 && opRes.NWords == 0 && opRes.NChars == 0 {
			continue
		}
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
	fmt.Print(finalResult)
	// return finalResult
}

func CalculateResult(flagOperations FlagOperations, filesToProcess []string, command string) {
	var finalOperationResult []OperationResults
	for _, file := range filesToProcess {
		operationResult, error := countOperation(flagOperations, file, command)
		if error != nil {
			fmt.Print(error.Error())
			continue
		}
		generateOutput([]OperationResults{operationResult}, flagOperations)
		finalOperationResult = append(finalOperationResult, operationResult)
	}

	// When processing for multiple files print the Total count as well
	if len(filesToProcess) > 1 {
		var totalResult OperationResults
		for _, opRes := range finalOperationResult {
			totalResult.NChars += opRes.NChars
			totalResult.NWords += opRes.NWords
			totalResult.NLines += opRes.NLines
		}
		totalResult.Filename += "total"
		generateOutput([]OperationResults{totalResult}, flagOperations)
	}
}
