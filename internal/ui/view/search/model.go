package search

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pauloo27/tuner/internal/ui/commands"
)

type model struct {
	query string
}

func NewModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	case commands.SearchCommand:
		m.query = msg.Query
	}

	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n",
		textStyle.Render(fmt.Sprintf("Searching for %s", m.query)),
	) + "\n"
}
