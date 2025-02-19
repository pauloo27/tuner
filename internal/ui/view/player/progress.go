package player

import (
	"log/slog"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pauloo27/tuner/internal/providers"
	"github.com/pauloo27/tuner/internal/ui/view"
)

type Progress struct {
	duration time.Duration
	position time.Duration
	percent  float64
}

func (Progress) ForwardTo() view.ViewName {
	return view.PlayerViewName
}

func fetchProgressNow() tea.Msg {
	var err error
	var p Progress

	p.duration, err = providers.Player.GetDuration()
	if err != nil {
		// while the file is loading, we might get this error, so just return empty progress
		if providers.Player.IsErrPropertyUnavailable(err) {
			return Progress{}
		}
		slog.Error("Failed to fetch duration", "err", err)
		return errMsg(err)
	}

	p.position, err = providers.Player.GetPosition()
	if err != nil {
		if providers.Player.IsErrPropertyUnavailable(err) {
			return Progress{}
		}
		slog.Error("Failed to fetch position", "err", err)
		return errMsg(err)
	}

	p.percent = float64(p.position) / float64(p.duration)

	return p
}

func fetchProgressTick() tea.Msg {
	time.Sleep(100 * time.Millisecond)
	return fetchProgressNow()
}
