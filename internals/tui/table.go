// NOTE: Bubbletea follows elm architecture
package tui

import (
	"fmt"

	_ "github.com/charmbracelet/bubbles"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

// Model
type Model struct {
	table table.Model
}

func NewModel(msg string) Model {
	return Model{
		
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

// Update
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	fmt.Println(msg)
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c":
			return m, tea.Quit
		}

	}
	return m, nil
}

func (m Model) FetchFromAPI() 

// View
func (m Model) View() string {
	return fmt.Sprintf("Hey there! ")
}
