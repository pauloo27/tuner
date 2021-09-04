package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const VERSION = "0.0.4-dev"

func main() {
	flex := tview.NewFlex()

	flex.SetBackgroundColor(tcell.ColorDefault)

	textView := tview.NewTextView().SetText("Tuner - " + VERSION)
	textView.SetBackgroundColor(tcell.ColorDefault)

	flex.AddItem(textView, 0, 1, true)

	if err := tview.NewApplication().SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
