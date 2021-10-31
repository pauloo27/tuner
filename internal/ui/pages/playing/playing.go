package playing

import (
	"github.com/Pauloo27/tuner/internal/providers/player"
	"github.com/Pauloo27/tuner/internal/providers/source"
	"github.com/Pauloo27/tuner/internal/ui"
	"github.com/Pauloo27/tuner/internal/ui/components/progress"
	"github.com/Pauloo27/tuner/internal/ui/components/progress/style"
	"github.com/Pauloo27/tuner/internal/utils"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	songProgress *progress.ProgressBar
	label        *tview.TextView
)

func onStart(params ...interface{}) {
	result := params[0].(*source.SearchResult)
	label.SetText(utils.Fmt("%s - %s", result.Artist, result.Title))
	go func() {
		err := player.Play(result)
		if err != nil {
			ui.App.QueueUpdateDraw(func() {
				label.SetText("Something went wrong...")
			})
		}
	}()
}

func init() {
	container := tview.NewGrid()
	container.SetColumns(0)
	container.SetRows(3, 1, 1)
	container.SetBackgroundColor(tcell.ColorDefault)

	label = tview.NewTextView()
	label.SetTextAlign(tview.AlignCenter)
	container.AddItem(label, 1, 0, 1, 1, 0, 0, false)

	songProgress = progress.NewProgressBar(style.NewSimpleBar("▏", "▎", "▍", "▌", "▋", "▊", "▉", "█"))
	songProgress.SetTextColor(tcell.ColorBlue)

	container.AddItem(songProgress, 2, 0, 1, 1, 0, 0, false)

	ui.RegisterPage(&ui.Page{
		Name: "playing", Container: container,
		OnStart: onStart,
	})
}