package playing

import (
	"github.com/Pauloo27/tuner/internal/providers/player"
	"github.com/Pauloo27/tuner/internal/providers/search"
	"github.com/Pauloo27/tuner/internal/ui"
	"github.com/Pauloo27/tuner/internal/utils"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var label *tview.TextView

func onStart(params ...interface{}) {
	result := params[0].(*search.SearchResult)
	label.SetText(utils.Fmt("%s - %s", result.Artist, result.Title))
	go func() {
		err := player.Play(result.URL)
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
	container.SetRows(1, 0)
	container.SetBackgroundColor(tcell.ColorDefault)

	label = tview.NewTextView()
	label.SetTextAlign(tview.AlignCenter)

	container.AddItem(label, 0, 0, 1, 1, 0, 0, false)

	ui.RegisterPage(&ui.Page{
		Name: "playing", Container: container,
		OnStart: onStart,
	})
}
