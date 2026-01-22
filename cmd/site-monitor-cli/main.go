package main

import "github.com/pradeept/site-monitor-cli/internals/logger"

func main() {
	// configure custom logger
	log := logger.Logger()

	// configure commands
	ConfigCommands()

	log.Print("Let's gooo")
}
