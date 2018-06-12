package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/resumic/cli/helper"
)

func main() {

	// Verify that a subcommand has been provided
	if len(os.Args) < 2 {
		fmt.Println("subcommand is required")
		os.Exit(1)
	}

	// Switch on the sub command
	switch os.Args[1] {
	case "init":
		//init.Parse(os.Args[2:])
		helper.InitResume()
		os.Exit(1)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
