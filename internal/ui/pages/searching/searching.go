package searching

import (
	"fmt"
	"log/slog"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/pauloo27/tuner/internal/providers"
	"github.com/pauloo27/tuner/internal/providers/source"
	"github.com/pauloo27/tuner/internal/ui"
	"github.com/pauloo27/tuner/internal/ui/core"
	"github.com/rivo/tview"
)

var resultList *tview.List
var label *tview.TextView

func init() {
	container := tview.NewGrid()
	container.SetColumns(0)
	container.SetRows(1, 0)

	label = tview.NewTextView()
	label.SetTextAlign(tview.AlignCenter)

	resultList = tview.NewList()
	resultList.SetSelectedBackgroundColor(tcell.ColorBlack)
	resultList.SetShortcutColor(tcell.ColorWhite)
	resultList.ShowSecondaryText(false)

	// vim-like keybinds (k and j navigation)
	resultList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'k':
			resultList.SetCurrentItem(resultList.GetCurrentItem() - 1)
		case 'j':
			item := resultList.GetCurrentItem() + 1
			if item >= resultList.GetItemCount() {
				item = 0
			}
			resultList.SetCurrentItem(item)
		}
		return event
	})

	container.AddItem(label, 0, 0, 1, 1, 0, 0, false)
	container.AddItem(resultList, 1, 0, 1, 1, 0, 0, true)

	ui.RegisterPage(&ui.Page{
		Name: "searching", Container: container,
		OnStart: onStart,
	})
}

func onStart(params ...interface{}) {
	searchQuery := params[0].(string)
	resultList.Clear()
	label.SetText(fmt.Sprintf("Searching for %s...", searchQuery))
	go func() {
		results, err := searchInAll(searchQuery)
		ui.App.QueueUpdateDraw(func() {
			if err != nil {
				slog.Error("Failed to search", "err", err)
				label.SetText("Something went wrong =(")
			}
			label.SetText(fmt.Sprintf("Results for %s:", searchQuery))
			for i, result := range results {
				// limit results to 10
				if i == 10 {
					break
				}
				shortcut := strconv.Itoa(i + 1)
				details := core.FmtEscaping(
					"[green]%s [white]from [green]%s [white]- %s", result.Title, result.Artist, result.Length,
				)

				currentResult := result

				resultList.AddItem(
					details, "", rune(shortcut[len(shortcut)-1]), func() {
						ui.SwitchPage("playing", currentResult)
					},
				)
			}
			resultList.AddItem("Cancel", "Press c to cancel", 'c', func() {
				ui.SwitchPage("home")
			})
		})
	}()
}

func searchInAll(query string) (results []source.SearchResult, err error) {
	for _, provider := range providers.Sources {
		r, err := provider.SearchFor(query)
		if err != nil {
			return nil, err
		}
		results = append(results, r...)
	}
	return
}
