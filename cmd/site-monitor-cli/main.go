package main

import (
	"os"
	"path"

	"github.com/pradeept/site-monitor-cli/internals/logger"
	"github.com/pradeept/site-monitor-cli/internals/store"
)

func main() {
	// configure custom logger
	log := logger.Logger()

	// configure commands
	ConfigCommands()

	// initialize store
	pwd, err := os.Getwd()
	log.Println(pwd)
	if err != nil {
		log.Fatal("Falied to fetch pwd")
		os.Exit(1)
	}

	db, err := store.NewStore(path.Join(pwd, "internals", "store", "app.db"))
	if err != nil {
		log.Println(err)
	}
	log.Println(db)
	log.Println("[Success] Store initialized")

	log.Print("Let's gooo")
}
