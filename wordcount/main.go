package main

import (
	"fmt"
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
	var executeOperations operations.FlagOperations
	var filesToProcess []string

	executeOperations, filesToProcess = parseArguments(cmdArgs[1:])
	operationResults := operations.CountOperations(executeOperations, filesToProcess, cmdArgs[0])
	// wordCountOutput := operations.GenerateOutput(operationResults, executeOperations)
	fmt.Print(operations.GenerateOutput(operationResults, executeOperations))
	if len(filesToProcess) > 1 {
		var totalResult operations.OperationResults
		for _, opRes := range operationResults {
			totalResult.NChars += opRes.NChars
			totalResult.NWords += opRes.NWords
			totalResult.NLines += opRes.NLines
		}
		totalResult.Filename += "Total"
		fmt.Print(operations.GenerateOutput([]operations.OperationResults{totalResult}, executeOperations))
	}
}
