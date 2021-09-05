package ui

import (
	"github.com/Pauloo27/tuner/internal/utils"
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

func SetFocus(component tview.Primitive) {
	app.SetFocus(component)
}

func SwitchPage(pageName string, params ...interface{}) {
	page, found := pageMap[pageName]
	pages.SwitchToPage(pageName)
	if !found {
		panic(utils.Fmt("Page %s not found", pageName))
	}
	if page.OnStart != nil {
		page.OnStart(params...)
	}
}
