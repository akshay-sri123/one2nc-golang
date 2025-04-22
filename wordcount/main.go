package main

import (
	"one2n/wordcount/operations"
	"os"
)

func setFlagOperations(passedOperations []string) operations.FlagOperations {
	var executeOperations operations.FlagOperations
	if len(passedOperations) == 0 {
		executeOperations.CWords = true
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
		command = cmdArgs[0]
		filename = cmdArgs[len(cmdArgs)-1]
		executeOperations = setFlagOperations(cmdArgs[1:])
	} else if len(cmdArgs) == 2 {
		command = cmdArgs[0]
		executeOperations = setFlagOperations(cmdArgs[1:])
		filename = cmdArgs[1]
	} else {
		command = cmdArgs[0]
		executeOperations = setFlagOperations(cmdArgs[1:])
		filename = cmdArgs[2]
	}

	operations.CheckFile(filename, command)
	operations.CountOperations(filename, executeOperations)
}
