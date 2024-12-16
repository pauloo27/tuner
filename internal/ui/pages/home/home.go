package home

import (
	"fmt"
	"log/slog"

	"github.com/gdamore/tcell/v2"
	"github.com/pauloo27/tuner/internal/core"
	"github.com/pauloo27/tuner/internal/ui"
	"github.com/pauloo27/tuner/internal/ui/theme"
	"github.com/rivo/tview"
)

type homePage struct {
	container   *tview.Grid
	searchInput *tview.InputField
}

func NewHomePage() *homePage {
	return &homePage{}
}

var _ ui.Page = &homePage{}

func (h *homePage) Init() error {
	h.container = tview.NewGrid().SetColumns(0).SetRows(1, 1, 0)
	pageTheme := theme.HomePageTheme

	label := tview.NewTextView().
		SetText(
			fmt.Sprintf("Tuner - %s", core.Version),
		).
		SetTextColor(pageTheme.TitleText)

	// FIXME: there must be a better way...
	emptySpace := tview.NewTextView()

	h.searchInput = tview.NewInputField().
		SetLabel(" > Search: ").
		SetLabelColor(pageTheme.SearchLabelText).
		SetFieldTextColor(pageTheme.SearchField.FG).
		SetFieldBackgroundColor(pageTheme.SearchField.BG)

	h.searchInput.SetDoneFunc(func(tcell.Key) {
		query := h.searchInput.GetText()
		if query == "" {
			return
		}
		slog.Info("Going to searching for", "query", query)
		ui.SwitchPage(ui.SearchingPageName, query)
	})

	h.container.AddItem(label, 0, 0, 1, 1, 0, 0, false).
		// can i do this using gap without the bg color? maybe something else? idk...
		AddItem(emptySpace, 1, 0, 1, 1, 0, 0, false).
		AddItem(h.searchInput, 2, 0, 1, 1, 0, 0, true)

	return nil
}

func (h *homePage) Name() ui.PageName {
	return ui.HomePageName
}

func (h *homePage) Open(...any) error {
	h.searchInput.SetText("")
	return nil
}

func (h *homePage) Container() tview.Primitive {
	return h.container
}
