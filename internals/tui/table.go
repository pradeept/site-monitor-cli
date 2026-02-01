package tui

import (
	_ "github.com/charmbracelet/bubbles"
	tea "github.com/charmbracelet/bubbletea"
)

// Model
type Model struct {
	SI       int
	SiteName string
	SiteId   int
}

func NewModel() *Model {
	return &Model{}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

// View

// Update
