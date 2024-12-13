package ui

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/pauloo27/tuner/internal/ui/theme"
	"github.com/rivo/tview"
)

const (
	defaultPageName = HomePageName
)

var (
	pagesByName = make(map[PageName]Page)
	tviewPages  = tview.NewPages()
	App         *tview.Application
)

func Setup() {
	theme.ResetTviewTheme()
}

func RegisterPages(pages ...Page) error {
	for _, page := range pages {
		if err := page.Init(); err != nil {
			return fmt.Errorf("failed to init page %s: %w", page.Name(), err)
		}
		pagesByName[page.Name()] = page
		tviewPages.AddPage(string(page.Name()), page.Container(), true, false)
	}
	return nil
}

func SwitchPage(pageName PageName, params ...any) {
	page, found := pagesByName[pageName]
	tviewPages.SwitchToPage(string(pageName))
	if !found {
		slog.Error("Page not found", "pageName", pageName)
		// nice
		os.Exit(69)
	}
	err := page.Open(params...)
	if err != nil {
		slog.Info("Failed to switch page", "page", page.Name(), "err", err)
		os.Exit(1)
	}
}

func StartTUI() error {
	App = tview.NewApplication()
	SwitchPage(defaultPageName)
	return App.SetRoot(tviewPages, true).Run()
}
