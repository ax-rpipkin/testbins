package main

import (
	"fmt"
	"os"
)

func main() {
	outputs := []string{
		"Automox Required Software Test Script\n\n",
		"This script will exit with success.\n\n",
		"To exit with failure, pass more than one argument\n",
	}

	errorouts := []string{
		"ERROR: More than one argument were provided",
	}

	for _, v := range outputs {
		fmt.Printf(v)
	}
	if len(os.Args) > 1 {
		for _, v := range errorouts {
			panic(v)
		}
	}
}
