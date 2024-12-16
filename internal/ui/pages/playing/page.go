package playing

import (
	"fmt"
	"log/slog"

	"github.com/gdamore/tcell/v2"
	"github.com/pauloo27/tuner/internal/providers"
	"github.com/pauloo27/tuner/internal/providers/source"
	"github.com/pauloo27/tuner/internal/ui"
	"github.com/pauloo27/tuner/internal/ui/core"
	"github.com/rivo/tview"
)

type playingPage struct {
	container    *tview.Flex
	result       source.SearchResult
	songLabel    *tview.TextView
	volumeLabel  *tview.TextView
	inputHandler map[rune]InputHandler
}

var _ ui.Page = &playingPage{}

func NewPlayingPage() *playingPage {
	return &playingPage{}
}

func (p *playingPage) Container() tview.Primitive {
	return p.container
}

func (p *playingPage) Init() error {
	p.container = tview.NewFlex().SetDirection(tview.FlexRow)
	p.songLabel = tview.NewTextView()
	p.volumeLabel = tview.NewTextView()
	p.container.AddItem(p.songLabel, 1, 1, true)
	p.container.AddItem(p.volumeLabel, 1, 1, false)
	p.container.AddItem(tview.NewTextView(), 0, 1, false)

	p.songLabel.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		p.handleInput(event.Rune())
		return event
	})

	p.registerInputHandlers()
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
	p.result = result
	p.songLabel.SetText(p.buildSongLabel(core.IconLoading))

	err := p.updateVolumeLabel()
	if err != nil {
		return err
	}

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
	isPaused, err := providers.Player.IsPaused()
	if err != nil {
		slog.Info("Failed to get paused status", "err", err)
		return
	}
	ui.App.QueueUpdateDraw(func() {
		if isPaused {
			p.songLabel.SetText(p.buildSongLabel(core.IconPaused))
		} else {
			p.songLabel.SetText(p.buildSongLabel(core.IconPlaying))
		}
	})
}
