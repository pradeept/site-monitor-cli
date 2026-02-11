package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pradeept/site-monitor-cli/internals/logger"
	"github.com/pradeept/site-monitor-cli/internals/store"
)

func ConfigCommands(s store.Store) {
	// sub-command add
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	siteName := addCmd.String("name", "", "Specify the name for website (ex: Google)")
	siteURL := addCmd.String("url", "", "Specify the url (ex: https://www.google.com)")
	reqTime := addCmd.Int64("time", 1, "Specify the time to call in seconds (ex: 2)")

	// sub-command remove
	removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	removeSiteName := removeCmd.String("website", "", "Specify the url (ex: https://www.google.com)")

	// USE CHARM

	// no arguments are passed
	if len(os.Args) < 2 {
		fmt.Println("Expected an argument, passed 0")
		// os.Exit(1)
		return
	}

	switch os.Args[1] {

	case "add":
		// parse flags for add sub-command
		addCmd.Parse(os.Args[2:])

		fmt.Println("[Add] Website: ", *siteName)
		if err := s.InsertSite(&store.Site{
			Id:          100,
			SiteName:    *siteName,
			SiteUrl:     *siteURL,
			RequestTime: *reqTime,
		}); err != nil {
			logger.Logger().Fatal("[Error] inserting the website: ", err)
			return
		}
		fmt.Println("Request Time: ", *reqTime)

	case "remove":
		// parse flags for remove sub-command
		removeCmd.Parse(os.Args[2:])
		fmt.Println("[Remove] Website: ", *removeSiteName)

		// default:
		// 	fmt.Println("Expected arguments, passed 0")
		// 	os.Exit(1)
	}
}
