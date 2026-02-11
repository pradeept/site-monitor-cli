// NOTE: Bubbletea follows elm architecture
package tui

import (
	"strconv"

	_ "github.com/charmbracelet/bubbles"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pradeept/site-monitor-cli/internals/logger"
	"github.com/pradeept/site-monitor-cli/internals/store"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

// Model
type Model struct {
	table table.Model
}

func NewModel(st store.Store) Model {
	log := logger.Logger()
	siteData, err := st.ListSites()
	if err != nil {
		log.Panic("Error occured while fetching sites: \n%v", err)
	}

	cols := []table.Column{
		{Title: "ID", Width: 15},
		{Title: "Site Name", Width: 25},
		{Title: "URL", Width: 50},
		{Title: "Request Time", Width: 25},
	}

	rows := []table.Row{}
	for _, site := range siteData {
		ss := table.Row{strconv.Itoa(site.Id), site.SiteName, site.SiteUrl, strconv.FormatInt(site.RequestTime, 10)}
		rows = append(rows, ss)
	}

	t := table.New(table.WithColumns(cols), table.WithFocused(true), table.WithHeight(7), table.WithRows(rows))
	t.Focus()
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
	return Model{table: t}
}

func (m Model) Init() tea.Cmd {
	return nil
}

// Update
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

// view
func (m Model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}
