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

var (
	totalLines, totalWords, totalChars int
)

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

func readTextFromFile(filename string, textChan chan<- string) {

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Print(err.Error())
	}

	textChan <- string(data)

	// return string(data), nil
}

func countOperation(executeOperations FlagOperations, textChan <-chan string) (OperationResults, error) {
	var operationResult OperationResults
	text := <-textChan
	// fmt.Print(text)

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

	return operationResult, nil
}

func (operationResult OperationResults) generateOutput(executeOperations FlagOperations) string {
	var finalResult string
	var output string
	if executeOperations.CLines {
		output += fmt.Sprintf("%8d ", operationResult.NLines)
	}

	if executeOperations.CWords {
		output += fmt.Sprintf("%8d ", operationResult.NWords)
	}

	if executeOperations.CChars {
		output += fmt.Sprintf("%8d ", operationResult.NChars)
	}

	if !executeOperations.CLines && !executeOperations.CWords && !executeOperations.CChars {
		output += fmt.Sprintf("%8d ", operationResult.NLines)
		output += fmt.Sprintf("%8d ", operationResult.NWords)
		output += fmt.Sprintf("%8d ", operationResult.NChars)
	}
	output += fmt.Sprint(" " + operationResult.Filename)
	finalResult += output + "\n"

	return finalResult
}

func CalculateResult(flagOperations FlagOperations, filesToProcess []string, command string) {
	textChan := make(chan string)
	defer close(textChan)

	for _, filename := range filesToProcess {
		err := checkFile(filename, command)
		if err != nil {
			fmt.Print(err.Error())
			continue
		}

		// text, err := readTextFromFile(file)
		// if err != nil {
		// 	fmt.Print(err.Error())
		// 	continue
		// }

		go readTextFromFile(filename, textChan)

		operationResult, err := countOperation(flagOperations, textChan)
		operationResult.Filename = filename

		totalLines += operationResult.NLines
		totalWords += operationResult.NWords
		totalChars += operationResult.NChars

		if err != nil {
			fmt.Print(err.Error())
		}

		fmt.Print(operationResult.generateOutput(flagOperations))
	}

	// When processing for multiple files print the Total count as well
	if len(filesToProcess) > 1 {
		var totalResult OperationResults
		totalResult.NChars = totalChars
		totalResult.NWords = totalWords
		totalResult.NLines = totalLines
		totalResult.Filename = "total"
		fmt.Print(totalResult.generateOutput(flagOperations))
	}
}
