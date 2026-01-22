package main

import (
	"flag"
	"fmt"
	"os"
)

func ConfigCommands() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addWebsite := addCmd.String("website", "", "Specify the url (ex: https://www.google.com)")
	callTime := addCmd.String("time", "", "Specify the time to call in seconds (ex: 2)")

	removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	removeWebsite := removeCmd.String("website", "", "Specify the url (ex: https://www.google.com)")

	if len(os.Args) < 2 {

		fmt.Println("Expected an argument, passed 0")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "add":
		addCmd.Parse(os.Args[2:])
		fmt.Println("[Add] Website: ", *addWebsite)
		fmt.Println("Time: ", *callTime)

	case "remove":
		removeCmd.Parse(os.Args[2:])
		fmt.Println("[Remove] Website: ", *removeWebsite)

	default:
		fmt.Println("Expected arguments, passed 0")
		os.Exit(1)
	}
}
