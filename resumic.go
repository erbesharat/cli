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
		helper.InitResume()
		os.Exit(1)
	case "serve":
		helper.ResumeServer(os.Args[1:])
		os.Exit(1)
	case "theme":
		if os.Args[2] == "get" {
			helper.GetTheme(os.Args[3])
		} else if os.Args[2] == "list" {
			helper.ListTheme()
		}
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
