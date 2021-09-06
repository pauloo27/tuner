package searching

import (
	"strconv"

	"github.com/Pauloo27/tuner/internal/providers/search"
	"github.com/Pauloo27/tuner/internal/ui"
	"github.com/Pauloo27/tuner/internal/utils"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var resultList *tview.List
var label *tview.TextView

func onStart(params ...interface{}) {
	searchQuery := params[0].(string)
	resultList.Clear()
	label.SetText(utils.Fmt("Searching for %s...", searchQuery))
	go func() {
		results, err := search.SearchInAll(searchQuery)
		ui.App.QueueUpdateDraw(func() {
			if err != nil {
				label.SetText("Something went wrong =(")
			}
			label.SetText(utils.Fmt("Results for %s:", searchQuery))
			for i, result := range results {
				// limit results to 10
				if i == 10 {
					break
				}
				shortcut := strconv.Itoa(i + 1)
				details := utils.Fmt("By [white]%s [green] - [white]%s[green]", result.Artist, result.Length)

				currentResult := result

				// TODO: escape colors?
				resultList.AddItem(
					result.Title, details, rune(shortcut[len(shortcut)-1]), func() {
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

func init() {
	container := tview.NewGrid()
	container.SetColumns(0)
	container.SetRows(1, 0)
	container.SetBackgroundColor(tcell.ColorDefault)

	label = tview.NewTextView()
	label.SetTextAlign(tview.AlignCenter)

	resultList = tview.NewList()
	resultList.SetSelectedBackgroundColor(tcell.ColorGreen)
	resultList.SetSelectedTextColor(tcell.ColorBlack)
	// vim-like keybinds (k and j navigation)
	resultList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'k':
			resultList.SetCurrentItem(resultList.GetCurrentItem() - 1)
		case 'j':
			resultList.SetCurrentItem(resultList.GetCurrentItem() + 1)
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
