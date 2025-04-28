package operations

import (
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
	filename string
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

func CheckFile(filesToProcess []string, command string) {
	for _, filename := range filesToProcess {
		if !checkIfFileExists(filename) {
			fmt.Fprintf(os.Stderr, "%s: %s: read: No such file or directory\n", command, filename)
			os.Exit(1)
		}

		if !checkIfFileOrDir(filename) {
			fmt.Fprintf(os.Stderr, "%s: %s: open: Is a directory\n", command, filename)
			os.Exit(1)
		}

		if !checkFilePermissions(filename) {
			fmt.Fprintf(os.Stderr, "%s: %s: open: Permission denied\n", command, filename)
			os.Exit(1)
		}
	}
}

func readTextFromFile(filename string) string {
	data, err := os.ReadFile(filename)
	check(err)

	return string(data)
}

func CountOperations(executeOperations FlagOperations, filesToProcess []string) []OperationResults {
	var operationResultsList []OperationResults
	for _, file := range filesToProcess {
		var operationResults OperationResults
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

		operationResults.filename = file
		operationResultsList = append(operationResultsList, operationResults)
	}
	return operationResultsList
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func PrintResults(operationResults []OperationResults, executeOperations FlagOperations) {
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
		output += fmt.Sprint(" " + opRes.filename)
		fmt.Println(output)
	}
}
