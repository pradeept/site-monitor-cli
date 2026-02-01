package main

import (
	"flag"
	"fmt"
	"os"
)

func ConfigCommands() {
	// sub-command add
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addWebsite := addCmd.String("website", "", "Specify the url (ex: https://www.google.com)")
	callTime := addCmd.String("time", "", "Specify the time to call in seconds (ex: 2)")

	// sub-command remove
	removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	removeWebsite := removeCmd.String("website", "", "Specify the url (ex: https://www.google.com)")

	// sub-command show to list a table of site and requests

	// USE CHARM

	// no arguments are passed
	if len(os.Args) < 2 {
		fmt.Println("Expected an argument, passed 0")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "add":
		// parse flags for add sub-command
		addCmd.Parse(os.Args[2:])
		fmt.Println("[Add] Website: ", *addWebsite)
		fmt.Println("Time: ", *callTime)

	case "remove":
		// parse flags for remove sub-command
		removeCmd.Parse(os.Args[2:])
		fmt.Println("[Remove] Website: ", *removeWebsite)
	
	case "show":


	default:
		fmt.Println("Expected arguments, passed 0")
		os.Exit(1)
	}
}
