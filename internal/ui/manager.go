package ui

import (
	"log/slog"
	"os"

	"github.com/rivo/tview"
)

var (
	App     *tview.Application
	pages   = tview.NewPages()
	pageMap = make(map[string]*Page)
)

func init() {
	tview.Styles = *GetTheme()
}

func GetTheme() *tview.Theme {
	// TODO: handle "theming"?
	return defaultTheme
}

func RegisterPage(page *Page) {
	pageMap[page.Name] = page
	pages.AddPage(page.Name, page.Container, true, false)
}

func StartApp(defaultPageName string) error {
	App = tview.NewApplication()
	SwitchPage(defaultPageName)
	return App.SetRoot(pages, true).Run()
}

func SetFocus(component tview.Primitive) {
	App.SetFocus(component)
}

func SwitchPage(pageName string, params ...interface{}) {
	page, found := pageMap[pageName]
	pages.SwitchToPage(pageName)
	if !found {
		slog.Error("Page not found", "pageName", pageName)
		// nice
		os.Exit(69)
	}
	if page.OnStart != nil {
		page.OnStart(params...)
	}
}
