package ui

import (
	"github.com/rivo/tview"
)

var (
	app   *tview.Application
	pages = tview.NewPages()
)

func init() {
	tview.Styles = *GetTheme()
}

func GetTheme() *tview.Theme {
	// TODO: handle "theming"?
	return defaultTheme
}

func RegisterPage(name string, page tview.Primitive) {
	pages.AddPage(name, page, true, false)
}

func StartApp(page string) error {
	app = tview.NewApplication()
	pages.SwitchToPage(page)
	return app.SetRoot(pages, true).Run()
}
