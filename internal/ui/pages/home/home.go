package pages

import (
	"fmt"
	"log/slog"

	"github.com/gdamore/tcell/v2"
	"github.com/pauloo27/tuner/internal/core"
	"github.com/pauloo27/tuner/internal/ui"
	"github.com/rivo/tview"
)

func init() {
	container := tview.NewGrid().SetColumns(0).SetRows(1, 1, 0)

	label := tview.NewTextView().
		SetText(
			fmt.Sprintf("Tuner - %s", core.Version),
		).
		SetTextColor(ui.GetTheme().TitleColor)

	// FIXME: there must be a better way...
	emptySpace := tview.NewTextView()

	searchInput := tview.NewInputField().
		SetFieldBackgroundColor(ui.GetTheme().PrimitiveBackgroundColor).
		SetLabel(" > Search: ")

	searchInput.SetDoneFunc(func(tcell.Key) {
		query := searchInput.GetText()
		if query == "" {
			return
		}
		slog.Info("Going to searching for", "query", query)
		ui.SwitchPage("searching", query)
	})

	container.AddItem(label, 0, 0, 1, 1, 0, 0, false).
		// can i do this using gap without the bg color? maybe something else? idk...
		AddItem(emptySpace, 1, 0, 1, 1, 0, 0, false).
		AddItem(searchInput, 2, 0, 1, 1, 0, 0, true)

	ui.RegisterPage(&ui.Page{
		Name: "home", Container: container,
		OnStart: func(...interface{}) {
			searchInput.SetText("")
		},
	})
}
