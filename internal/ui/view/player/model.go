package player

import (
	"errors"
	"fmt"
	"log/slog"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pauloo27/tuner/internal/providers/player"
	"github.com/pauloo27/tuner/internal/providers/source"
	"github.com/pauloo27/tuner/internal/ui/uicore"
	"github.com/pauloo27/tuner/internal/ui/view"
)

type model struct {
	playing          source.SearchResult
	eventLoopStarted bool
	volume           float64
	isLoading        bool
	isPaused         bool
	err              error
}

func NewModel() model {
	return model{}
}

func (model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case InitialState:
		m.volume = msg.volume
		m.isPaused = msg.isPaused
	case Event:
		slog.Info("Got player event", "event", msg)
		switch msg.name {
		case player.PlayerEventPlay:
			m.isPaused = false
		case player.PlayerEventPause:
			m.isPaused = true
		case player.PlayerEventVolumeChanged:
			m.volume = msg.params[0].(float64)
		}
		cmds = append(cmds, fetchEvent)
	case errMsg:
		m.err = msg
		m.isLoading = false
	case loadedMsg:
		m.err = nil
		m.isLoading = false
	case tea.KeyMsg:
		switch msg.String() {
		case " ":
			cmds = append(cmds, togglePause)
		case "-":
			cmds = append(cmds, decreaseVolume)
		case "+", "=":
			cmds = append(cmds, increaseVolume)
		default:
			slog.Info("Unhandled key", "key", msg.String())
		}
	case view.GotoViewMsg:
		if !m.eventLoopStarted {
			listenToEvents()
			m.eventLoopStarted = true
			cmds = append(cmds, loadInitialState, fetchEvent)
		}

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
		cmds = append(cmds, play(result))
	}

	return m, tea.Batch(cmds...)
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
			primaryTextStyle.Render(fmt.Sprintf("%s...", m.playing.Title)),
		)
	}

	icon := uicore.IconPlaying
	if m.isPaused {
		icon = uicore.IconPaused
	}

	return fmt.Sprintf(
		"%s\n\n%s\n",
		primaryTextStyle.Render(fmt.Sprintf("%s | %s by %s", icon, m.playing.Title, m.playing.Artist)),
		secondaryTextStyle.Render(fmt.Sprintf("Volume: %.0f%%", m.volume)),
	)
}
