package searching

import (
	"github.com/Pauloo27/tuner/internal/ui"
	"github.com/Pauloo27/tuner/internal/utils"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func init() {
	container := tview.NewGrid()
	container.SetColumns(0)
	container.SetRows(1, 0)
	container.SetBackgroundColor(tcell.ColorDefault)

	label := tview.NewTextView()
	label.SetTextAlign(tview.AlignCenter)

	container.AddItem(label, 0, 0, 1, 1, 0, 0, false)
	//container.AddItem(searchInput, 1, 0, 1, 1, 0, 0, true)

	ui.RegisterPage(&ui.Page{
		Name: "searching", Container: container,
		OnStart: func(params ...interface{}) {
			searchQuery := params[0]
			label.SetText(utils.Fmt("Results for %s", searchQuery))
		},
	})
}
