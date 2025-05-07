package main

import (
	"one2n/wordcount/operations"
	"os"
)

func parseArguments(passedOperations []string) (operations.FlagOperations, []string) {
	var executeOperations operations.FlagOperations
	var filenames []string
	if len(passedOperations) == 0 {
		executeOperations.CLines = true
		executeOperations.CWords = true
		executeOperations.CChars = true
	} else {
		for _, operation := range passedOperations {
			if operation == "-l" {
				executeOperations.CLines = true
			} else if operation == "-w" {
				executeOperations.CWords = true
			} else if operation == "-c" {
				executeOperations.CChars = true
			} else {
				filenames = append(filenames, operation)
			}
		}
	}
	return executeOperations, filenames
}

func main() {
	cmdArgs := os.Args
	var flagOperations operations.FlagOperations
	var filesToProcess []string

	flagOperations, filesToProcess = parseArguments(cmdArgs[1:])
	operations.CalculateResult(flagOperations, filesToProcess, cmdArgs[0])
}
