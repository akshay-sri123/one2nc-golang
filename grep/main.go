package main

import (
	"bufio"
	"log"
	"one2n/grep/operations"
	"os"
)

func parseArguments(flagOperations operations.FlagOperations, passedOperations []string) operations.FlagOperations {
	flagOperations.FilterString = passedOperations[1]
	var stdIn string

	if len(passedOperations) == 2 {
		scanner := bufio.NewScanner(os.Stdin)

		for {
			scanner.Scan()
			line := scanner.Text()
			if len(line) == 0 {
				break
			}

			stdIn = stdIn + line + "\n"
		}
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		flagOperations.StdInput = stdIn

	} else {
		flagOperations.FilesToProcess = passedOperations[2]
	}
	return flagOperations
}

func main() {
	cmdArgs := os.Args
	var flagOperations operations.FlagOperations

	flagOperations = parseArguments(flagOperations, cmdArgs)

	operations.RunOperation(flagOperations, cmdArgs[0])
}
