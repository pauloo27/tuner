package pages

import (
	"github.com/Pauloo27/tuner/ui"
	"github.com/Pauloo27/tuner/version"
	"github.com/rivo/tview"
)

func init() {
	container := tview.NewGrid()
	container.SetBackgroundColor(ui.GetTheme().PrimitiveBackgroundColor)
	container.SetColumns(0)
	container.SetRows(1, 0)

	label := tview.NewTextView()
	label.SetBackgroundColor(ui.GetTheme().PrimitiveBackgroundColor)
	label.SetText("Tuner - " + version.Current)
	label.SetTextAlign(tview.AlignCenter)

	searchInput := tview.NewInputField()
	searchInput.SetBackgroundColor(ui.GetTheme().PrimitiveBackgroundColor)
	searchInput.SetFieldBackgroundColor(ui.GetTheme().PrimitiveBackgroundColor)
	searchInput.SetLabel("Search for: ")

	container.AddItem(label, 0, 0, 1, 1, 0, 0, false)
	container.AddItem(searchInput, 1, 0, 1, 1, 0, 0, true)

	ui.RegisterPage("home", container)
}
