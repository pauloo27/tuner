package ui

import "github.com/rivo/tview"

type PageName string

const (
	HomePageName      PageName = "HOME"
	SearchingPageName PageName = "SEARCHING"
	PlayingPageName   PageName = "PLAYING"
)

type Page interface {
	Name() PageName
	Init() error
	Container() tview.Primitive
	Open(...any) error
}
