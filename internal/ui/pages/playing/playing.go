package playing

import (
	"fmt"
	"log/slog"

	"github.com/gdamore/tcell/v2"
	"github.com/pauloo27/tuner/internal/providers"
	"github.com/pauloo27/tuner/internal/providers/player"
	"github.com/pauloo27/tuner/internal/providers/source"
	"github.com/pauloo27/tuner/internal/ui"
	"github.com/pauloo27/tuner/internal/ui/core"
	"github.com/rivo/tview"
)

type playingPage struct {
	container  *tview.Grid
	songLabel  *tview.TextView
	songStatus *tview.TextView
}

var _ ui.Page = &playingPage{}

func NewPlayingPage() *playingPage {
	return &playingPage{}
}

func (p *playingPage) Container() tview.Primitive {
	return p.container
}

func (p *playingPage) Init() error {
	p.container = tview.NewGrid().SetColumns(0).SetRows(1, -3)
	p.songLabel = tview.NewTextView()
	p.songStatus = tview.NewTextView()
	p.container.AddItem(p.songLabel, 0, 0, 1, 1, 0, 0, false)
	p.container.AddItem(p.songStatus, 1, 0, 1, 1, 0, 0, false)

	p.container.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		p.handleInput(event.Rune())
		return event
	})

	p.registerListeners()

	return nil
}

func (p *playingPage) Name() ui.PageName {
	return ui.PlayingPageName
}

func (p *playingPage) Open(params ...any) error {
	if len(params) != 1 {
		return fmt.Errorf("expected 1 parameter, got %d", len(params))
	}

	result, ok := params[0].(source.SearchResult)
	if !ok {
		return fmt.Errorf("parameter is not a SearchResult")
	}
	p.songLabel.SetText(fmt.Sprintf("%s - %s", result.Artist, result.Title))
	p.songStatus.SetText(core.IconLoading)

	go p.play(result)
	return nil
}

func (p *playingPage) play(result source.SearchResult) {
	err := providers.Player.Play(result)
	if err != nil {
		slog.Error("Failed to play song", "player", providers.Player.Name(), "err", err)
		ui.App.QueueUpdateDraw(func() {
			p.songLabel.SetText("Something went wrong...")
		})
	}
}

func (p *playingPage) registerListeners() {
	providers.Player.On(player.PlayerEventPause, func(...any) {
		ui.App.QueueUpdateDraw(func() {
			p.songStatus.SetText(core.IconPaused)
		})
	})

	providers.Player.On(player.PlayerEventPlay, func(...any) {
		ui.App.QueueUpdateDraw(func() {
			p.songStatus.SetText(core.IconPlaying)
		})
	})
}

func (p *playingPage) handleInput(key rune) {
	switch key {
	case ' ':
		err := providers.Player.TogglePause()
		if err != nil {
			slog.Error("Failed to toggle pause", "err", err)
		}
	default:
		slog.Info("Unhandled input", "key", key)
	}
}
