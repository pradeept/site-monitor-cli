/*
This file is used as a playground. Will be trashed later.
*/
package main

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	_ "github.com/mattn/go-sqlite3" //driver
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func main() {

	p := tea.NewProgram(
		InitialModel(),
	)
	if err := p.Start(); err != nil {
		panic(err)
	}
}

type Model struct {
	table table.Model
}

func InitialModel() Model {
	cols := []table.Column{
`		{Title: "SI No.", Width: 15},
`		{Title: "Website", Width: 25},
		{Title: "Avg. Success Rate", Width: 50},
	}

	rows := []table.Row{
		{"1", "https://www.google.com", "100%"},
		{"2", "https://pradeept.dev", "100%"},
		{"3", "https://youtube.com", "100%"},
		{"4", "https://facebook.com", "100%"},
	}
	t := table.New(table.WithColumns(cols), table.WithFocused(true), table.WithHeight(7), table.WithRows(rows))
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := Model{t}
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}
