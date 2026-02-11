package main

import (
	"os"
	"path"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pradeept/site-monitor-cli/internals/logger"
	"github.com/pradeept/site-monitor-cli/internals/store"
	"github.com/pradeept/site-monitor-cli/internals/tui"
)

func main() {
	// configure custom logger
	log := logger.Logger()

	// initialize store
	pwd, err := os.Getwd()

	if err != nil {
		log.Fatal("Falied to fetch pwd")
		os.Exit(1)
	}

	store, err := store.NewStore(path.Join(pwd, "internals", "store", "app.db"))
	if err != nil {
		log.Println(err)
	}
	log.Println("[Success] Store initialized")

	// configure commands
	ConfigCommands(*store)

	p := tea.NewProgram(
		tui.NewModel(*store),
	)
	if err = p.Start(); err != nil {
		panic(err)
	}
}
