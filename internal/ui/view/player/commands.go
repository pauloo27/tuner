package player

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pauloo27/tuner/internal/providers"
	"github.com/pauloo27/tuner/internal/providers/source"
)

type errMsg error

type loadedMsg struct{}

func play(result source.SearchResult) tea.Cmd {
	return func() tea.Msg {
		err := providers.Player.Play(result)

		if err != nil {
			return errMsg(err)
		}

		return loadedMsg{}
	}
}
