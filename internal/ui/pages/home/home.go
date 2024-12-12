package pages

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/pauloo27/tuner/internal/core"
	"github.com/pauloo27/tuner/internal/ui"
	"github.com/rivo/tview"
)

func init() {
	container := tview.NewGrid()
	container.SetColumns(0)
	container.SetRows(1, 0)

	label := tview.NewTextView()
	label.SetText(fmt.Sprintf("[green:black]Tuner - %s", core.Version))
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
