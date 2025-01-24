package home

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pauloo27/tuner/internal/core"
	"github.com/pauloo27/tuner/internal/ui/view/root"
)

type model struct {
	searchInput textinput.Model
}

func NewModel() model {
	ti := textinput.New()
	ti.Placeholder = "Search..."
	ti.Focus()
	ti.CharLimit = 100
	return model{
		searchInput: ti,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			return m, root.StartSearch(m.searchInput.Value())
		}
	}

	m.searchInput, cmd = m.searchInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		textStyle.Render(fmt.Sprintf("Tuner - %s", core.Version)),
		m.searchInput.View(),
	) + "\n"
}
