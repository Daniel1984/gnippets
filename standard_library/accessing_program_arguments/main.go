package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	// This call will print
	// all command line arguments.
	fmt.Println(args)
	// The first argument, zero item from slice,
	// is the name of the called binary.
	programName := args[0]
	fmt.Printf("The binary name is: %s \n", programName)
	// The rest of the arguments could be obtained
	// by omitting the first argument.
	otherArgs := args[1:]
	fmt.Println(otherArgs)

	for idx, arg := range otherArgs {
		fmt.Printf("Arg %d = %s \n", idx, arg)
	}
}
