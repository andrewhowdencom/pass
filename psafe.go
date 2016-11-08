package main

import "fmt"
import "os"

func main() {
	var args = parseArgs()

	fmt.Printf(args)
}

func parseArgs() string {
	const defaultArg string = "help"
	var arg string

	if len(os.Args) > 1 {
		arg = os.Args[1]
	} else {
		arg = defaultArg
	}

	return arg
}
