package playing

import (
	"fmt"
	"log/slog"

	"github.com/pauloo27/tuner/internal/providers"
	"github.com/pauloo27/tuner/internal/providers/source"
	"github.com/pauloo27/tuner/internal/ui"
	"github.com/rivo/tview"
)

type playingPage struct {
	container *tview.Grid
	label     *tview.TextView
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
	p.label = tview.NewTextView()
	p.container.AddItem(p.label, 0, 0, 1, 1, 0, 0, false)
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
	p.label.SetText(fmt.Sprintf("%s - %s", result.Artist, result.Title))

	go p.play(result)
	return nil
}

func (p *playingPage) play(result source.SearchResult) {
	err := providers.Player.Play(result)
	if err != nil {
		slog.Error("Failed to play song", "player", providers.Player.GetName(), "err", err)
		ui.App.QueueUpdateDraw(func() {
			p.label.SetText("Something went wrong...")
		})
	}
}
