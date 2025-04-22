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

func CheckFile(filename string, command string) {
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

func readTextFromFile(filename string) string {
	data, err := os.ReadFile(filename)
	check(err)

	return string(data)
}

func CountOperations(filename string, executeOperations FlagOperations) {
	text := readTextFromFile(filename)
	outputString := "%8d %s\n"
	if executeOperations.CLines {
		outputString = fmt.Sprintf("%8d %s\n", countLines(text), filename)
	} else if executeOperations.CWords {
		outputString = fmt.Sprintf("%8d %s\n", countWords(text), filename)
	} else if executeOperations.CChars {
		outputString = fmt.Sprintf("%8d %s\n", countCharacters(text), filename)
	} else {
		outputString = fmt.Sprintf("%8d %8d %8d %s\n", countLines(text), countWords(text), countCharacters(text), filename)
	}

	fmt.Fprint(os.Stdout, outputString)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
