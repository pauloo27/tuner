package player

import (
	"log/slog"

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

func togglePause() tea.Msg {
	err := providers.Player.TogglePause()
	if err != nil {
		slog.Error("Failed to toggle pause", "err", err)
		return errMsg(err)
	}
	return nil
}

func increaseVolume() tea.Msg {
	curVolume, err := providers.Player.GetVolume()
	if err != nil {
		slog.Error("Failed to get current volume", "err", err)
		return errMsg(err)
	}

	err = providers.Player.SetVolume(curVolume + 1)
	if err != nil {
		slog.Error("Failed to set volume", "err", err)
		return errMsg(err)
	}
	return nil
}

func decreaseVolume() tea.Msg {
	curVolume, err := providers.Player.GetVolume()
	if err != nil {
		slog.Error("Failed to get current volume", "err", err)
		return errMsg(err)
	}

	err = providers.Player.SetVolume(curVolume - 1)
	if err != nil {
		slog.Error("Failed to set volume", "err", err)
		return errMsg(err)
	}
	return nil
}

func fetchEvent() tea.Msg {
	event := <-eventCh
	return event
}
