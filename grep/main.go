package main

import (
	"one2n/grep/operations"
	"os"
)

func parseArguments(passedOperations []string) operations.FlagOperations {
	var flagOperations operations.FlagOperations
	flagOperations.FilterString = passedOperations[0]
	flagOperations.FilesToProcess = passedOperations[1]
	return flagOperations
}

func main() {
	cmdArgs := os.Args
	var flagOperations operations.FlagOperations

	flagOperations = parseArguments(cmdArgs[1:])
	operations.RunOperation(flagOperations, cmdArgs[0])
}
