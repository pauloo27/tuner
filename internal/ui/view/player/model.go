package player

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/charmbracelet/bubbles/progress"
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
	progress         progress.Model
	percent          float64
}

func NewModel() model {
	m := model{}
	m.progress = progress.New(progress.WithSolidFill(string(progressFill)), progress.WithoutPercentage())
	return m
}

func (model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case view.GotoViewMsg:
		if !m.eventLoopStarted {
			listenToEvents()
			m.eventLoopStarted = true
			cmds = append(cmds, fetchProgressNow, fetchInitialState, fetchEvent)
		}

		if len(msg.Params) != 1 {
			slog.Error("Expected one param in player view", "got", len(msg.Params))
			m.err = errors.New("invalid params count received")
			return m, nil
		}
		result, ok := msg.Params[0].(source.SearchResult)
		if !ok {
			slog.Error("Expected SearchResult param in player view")
			m.err = errors.New("invalid params type received")
			return m, nil
		}
		m.playing = result
		m.isLoading = true
		cmds = append(cmds, play(result))
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
	case Progress:
		slog.Info("Got progress", "p", msg)
		m.percent = msg.percent
		cmds = append(cmds, fetchProgressTick)
	case errMsg:
		m.err = msg
	case loadedMsg:
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
	case view.VisibleMsg:
		m.progress.Width = msg.Width
	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width
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
		"%s\n\n%s\n%s\n",
		m.progress.ViewAs(m.percent),
		primaryTextStyle.Render(fmt.Sprintf("%s | %s by %s", icon, m.playing.Title, m.playing.Artist)),
		secondaryTextStyle.Render(fmt.Sprintf("Volume: %.0f%%", m.volume)),
	)
}
