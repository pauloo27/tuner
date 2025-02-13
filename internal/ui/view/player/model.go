package player

import (
	"errors"
	"fmt"
	"log/slog"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pauloo27/tuner/internal/providers/source"
	"github.com/pauloo27/tuner/internal/ui/view"
)

type model struct {
	playing   source.SearchResult
	isLoading bool
	err       error
}

func NewModel() model {
	return model{}
}

func (model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case errMsg:
		m.err = msg
		m.isLoading = false
	case loadedMsg:
		m.err = nil
		m.isLoading = false
	case view.GotoViewMsg:
		if len(msg.Params) != 1 {
			slog.Info("Expected one param in player view", "got", len(msg.Params))
			m.err = errors.New("invalid params count received")
			return m, nil
		}
		result, ok := msg.Params[0].(source.SearchResult)
		if !ok {
			slog.Info("Expected SearchResult param in player view")
			m.err = errors.New("invalid params type received")
			return m, nil
		}
		m.playing = result
		m.isLoading = true
		cmd = play(result)
	}

	return m, cmd
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf(
			"%s\n",
			errorStyle.Render(m.err.Error()),
		)
	}

	if m.isLoading {
		// TODO: spinner
		return fmt.Sprintf(
			"%s\n",
			textStyle.Render(fmt.Sprintf("%s...", m.playing.Title)),
		)
	}

	return fmt.Sprintf(
		"%s\n",
		textStyle.Render(m.playing.Title),
	)
}
