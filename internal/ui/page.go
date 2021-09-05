package ui

import "github.com/rivo/tview"

type Page struct {
	Name      string
	Container tview.Primitive
	OnStart   func()
}
