package playing

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/pauloo27/tuner/internal/providers"
	"github.com/pauloo27/tuner/internal/providers/source"
	"github.com/pauloo27/tuner/internal/ui"
	"github.com/pauloo27/tuner/internal/ui/components/progress"
	"github.com/pauloo27/tuner/internal/ui/components/progress/style"
	"github.com/pauloo27/tuner/internal/ui/core"
	"github.com/pauloo27/tuner/internal/ui/theme"
	"github.com/rivo/tview"
)

type playingPage struct {
	container     *tview.Flex
	result        source.SearchResult
	progressBar   *progress.ProgressBar
	songLabel     *tview.TextView
	volumeLabel   *tview.TextView
	inputsHandler map[rune]InputHandler
}

var _ ui.Page = &playingPage{}

func NewPlayingPage() *playingPage {
	return &playingPage{}
}

func (p *playingPage) Container() tview.Primitive {
	return p.container
}

func (p *playingPage) Init() error {
	pageTheme := theme.PlayingPageTheme

	p.container = tview.NewFlex().SetDirection(tview.FlexRow)
	p.songLabel = tview.NewTextView().SetTextColor(pageTheme.SongInfoColor)
	p.volumeLabel = tview.NewTextView().SetTextColor(pageTheme.VolumeColor)
	p.progressBar = progress.NewProgressBar(style.NewSimpleBarWithBlocks())

	p.container.AddItem(p.progressBar, 1, 1, false)
	p.container.AddItem(p.songLabel, 1, 1, true)
	p.container.AddItem(p.volumeLabel, 2, 1, false)
	p.container.AddItem(tview.NewTextView(), 0, 1, false)

	p.songLabel.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		p.handleInput(event.Rune())
		return event
	})

	p.registerInputHandlers()
	p.registerListeners()
	p.startProgressLoop()

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

func (p *playingPage) startProgressLoop() {
	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			// skip if not on the playing page
			if !p.songLabel.HasFocus() {
				continue
			}

			duration, err := providers.Player.GetDuration()
			if err != nil {
				slog.Error("Failed to fetch player duration", "err", err)
				continue
			}

			position, err := providers.Player.GetPosition()
			if err != nil {
				slog.Error("Failed to fetch player position", "err", err)
				continue
			}

			ui.App.QueueUpdateDraw(func() {
				p.progressBar.SetRelativeProgress(float64(duration), float64(position))
			})
		}
	}()
}
