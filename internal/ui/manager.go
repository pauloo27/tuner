package ui

import (
	"github.com/rivo/tview"
)

var (
	app     *tview.Application
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
	app = tview.NewApplication()
	SwitchPage(defaultPageName)
	return app.SetRoot(pages, true).Run()
}

func SwitchPage(pageName string) {
	page, found := pageMap[pageName]
	pages.SwitchToPage(pageName)
	if found && page.OnStart != nil {
		page.OnStart()
	}
}
