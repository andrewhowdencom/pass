package main

import "fmt"
import "os"

/**
 * Program entrypoint.
 */
func main() {
	var args, isValid = parseArgs()
	result := "Args are valid"

	if isValid == false {
		result = "Args are not valid"
	}

	fmt.Printf(args + "\n")
	fmt.Printf(result + "\n")
}

/**
 * Take the command line arguments, and return a structure that the rest of the program can work with. Responsible for
 * handling any errors, such as no arguments or malformed arguments.
 */
func parseArgs() (string, bool) {
	const defaultArg string = "help"

	allowedArgs := []string{"help", "ls", "find", "show", "grep", "insert", "edit", "generate", "rm", "mv", "cp", "version"}
	arg := "help"
	isValid := false

	if len(os.Args) > 1 {
		arg = os.Args[1]
	}

	// Determine if the argument is even allowed.
	for _, compare := range allowedArgs {
		if compare == arg {
			isValid = true
		}
	}

	return arg, isValid
}
