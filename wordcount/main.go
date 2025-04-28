package main

import (
	"one2n/wordcount/operations"
	"os"
)

func setFlagOperations(passedOperations []string) (operations.FlagOperations, []string) {
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
	command := ""
	filesToProcess := []string{}
	var executeOperations operations.FlagOperations

	if len(cmdArgs) > 3 {
		// Passing ./wc -l -w -c filename1 filename2 filename3
		command = cmdArgs[0]
		executeOperations, filesToProcess = setFlagOperations(cmdArgs[1:])
	} else if len(cmdArgs) == 2 {
		// Passing ./wc filename
		command = cmdArgs[0]
		executeOperations, filesToProcess = setFlagOperations(cmdArgs[1:])
	} else {
		// Passing ./wc -l filename
		command = cmdArgs[0]
		executeOperations, filesToProcess = setFlagOperations(cmdArgs[1:])
	}

	operations.CheckFile(filesToProcess, command)
	operationResults := operations.CountOperations(executeOperations, filesToProcess)
	// fmt.Println(filesToProcess)
	// fmt.Println(operationResults)
	operations.PrintResults(operationResults, executeOperations)
}
