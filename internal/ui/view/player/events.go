package player

import (
	"log/slog"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pauloo27/tuner/internal/providers"
	"github.com/pauloo27/tuner/internal/providers/player"
	"github.com/pauloo27/tuner/internal/ui/view"
)

var (
	eventCh = make(chan Event) // should I buffer it? idk
)

type Event struct {
	name   player.PlayerEvent
	params []any
}

type InitialState struct {
	isPaused bool
	volume   float64
}

func (Event) ForwardTo() view.ViewName {
	return view.PlayerViewName
}

func fetchInitialState() tea.Msg {
	var err error
	var state InitialState

	state.isPaused, err = providers.Player.IsPaused()
	if err != nil {
		slog.Error("Failed to get initial pause status", "err", err)
		return errMsg(err)
	}

	state.volume, err = providers.Player.GetVolume()
	if err != nil {
		slog.Error("Failed to get initial volume", "err", err)
		return errMsg(err)
	}

	return state
}

func listenToEvents() {
	listenAndQueue := func(names ...player.PlayerEvent) {
		for _, name := range names {
			providers.Player.On(name, func(params ...any) {
				queueEvent(name, params...)
			})
		}
	}

	listenAndQueue(player.PlayerEventPlay, player.PlayerEventPause, player.PlayerEventVolumeChanged)
}

func queueEvent(name player.PlayerEvent, params ...any) {
	eventCh <- Event{name, params}
}
