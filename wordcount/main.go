package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
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

func check(e error) {
	if e != nil {
		panic(e)
	}
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

func checkFile(filename string, command string) {
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

func countOperations(filename string, operationFlag string) {
	text := readTextFromFile(filename)
	switch operationFlag {
	case "-l":
		fmt.Fprintf(os.Stdout, "%8d %s\n", countLines(text), filename)
	case "-w":
		fmt.Fprintf(os.Stdout, "%8d %s\n", countWords(text), filename)
	case "-c":
		fmt.Fprintf(os.Stdout, "%8d %s\n", countCharacters(text), filename)
	}
}

func main() {
	cmdArgs := os.Args
	command := cmdArgs[0]
	operationFlag := cmdArgs[1]
	filename := cmdArgs[2]

	if len(cmdArgs) != 3 || !slices.Contains([]string{"-l", "-c", "-w"}, operationFlag) {
		fmt.Fprintf(os.Stderr, "Incorrect usage.\n")
		os.Exit(1)
	}

	checkFile(filename, command)
	countOperations(filename, operationFlag)
}
