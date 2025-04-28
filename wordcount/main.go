package main

import (
	"one2n/wordcount/operations"
	"os"
)

func setFlagOperations(passedOperations []string) operations.FlagOperations {
	var executeOperations operations.FlagOperations
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
			}
		}
	}
	return executeOperations
}

func main() {
	cmdArgs := os.Args
	filename, command := "", ""
	var executeOperations operations.FlagOperations

	if len(cmdArgs) > 3 {
		// Passing ./wc -l -w -c filename
		command = cmdArgs[0]
		filename = cmdArgs[len(cmdArgs)-1]
		executeOperations = setFlagOperations(cmdArgs[1:])
	} else if len(cmdArgs) == 2 {
		// Passing ./wc filename
		command = cmdArgs[0]
		executeOperations = setFlagOperations([]string{})
		filename = cmdArgs[1]
	} else {
		// Passing ./wc -l filename
		command = cmdArgs[0]
		executeOperations = setFlagOperations([]string{cmdArgs[1]})
		filename = cmdArgs[2]
	}

	operations.CheckFile(filename, command)
	operationResults := operations.CountOperations(filename, executeOperations)
	operations.PrintResults(operationResults, executeOperations)
}
