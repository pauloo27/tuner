package pages

import (
	"github.com/Pauloo27/tuner/internal/ui"
	"github.com/Pauloo27/tuner/internal/utils"
	"github.com/Pauloo27/tuner/internal/version"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func init() {
	container := tview.NewGrid()
	container.SetColumns(0)
	container.SetRows(1, 0)

	label := tview.NewTextView()
	label.SetText(utils.Fmt("[green:black]Tuner - %s", version.Current))
	label.SetTextAlign(tview.AlignCenter).SetDynamicColors(true)

	searchInput := tview.NewInputField()
	searchInput.SetFieldBackgroundColor(ui.GetTheme().PrimitiveBackgroundColor)
	searchInput.SetLabel("[blue]Search for: ")
	searchInput.SetDoneFunc(func(tcell.Key) {
		ui.SwitchPage("searching", searchInput.GetText())
	})

	container.AddItem(label, 0, 0, 1, 1, 0, 0, false)
	container.AddItem(searchInput, 1, 0, 1, 1, 0, 0, true)

	ui.RegisterPage(&ui.Page{
		Name: "home", Container: container,
		OnStart: func(...interface{}) {
			searchInput.SetText("")
		},
	})
}
