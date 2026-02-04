package tui

import (
	"fmt"

	_ "github.com/charmbracelet/bubbles"
	tea "github.com/charmbracelet/bubbletea"
)

// Model
type Model struct {
	SI       int
	SiteName string
	SiteId   int
	Message  string
}

func NewModel(msg string) Model {
	return Model{
		Message: msg,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

// Update
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c":
			return m, tea.Quit
		}

	}
	return m, nil
}

// View
func (m Model) View() string {
	return fmt.Sprintf("Hey there! ")
}
