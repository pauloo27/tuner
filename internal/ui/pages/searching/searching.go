package searching

import (
	"fmt"
	"log/slog"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/pauloo27/tuner/internal/providers"
	"github.com/pauloo27/tuner/internal/providers/source"
	"github.com/pauloo27/tuner/internal/ui"
	"github.com/pauloo27/tuner/internal/ui/core"
	"github.com/rivo/tview"
)

type searchingPage struct {
	resultList *tview.List
	label      *tview.TextView
	container  *tview.Grid
}

var _ ui.Page = &searchingPage{}

func (s *searchingPage) Container() tview.Primitive {
	return s.container
}

func (s *searchingPage) Init() error {
	s.container = tview.NewGrid().SetColumns(0).SetRows(1, 0)
	s.label = tview.NewTextView()
	s.resultList = s.newResultList()

	s.container.AddItem(s.label, 0, 0, 1, 1, 0, 0, false)
	s.container.AddItem(s.resultList, 1, 0, 1, 1, 0, 0, true)
	return nil
}

func (s *searchingPage) Name() ui.PageName {
	return ui.SearchingPageName
}

func (s *searchingPage) Open(params ...any) error {
	if len(params) != 1 {
		return fmt.Errorf("expected 1 parameter, got %d", len(params))
	}
	searchQuery, ok := params[0].(string)
	if !ok {
		return fmt.Errorf("parameter is not a string")
	}
	s.resultList.Clear()
	go s.searchAndThenUpdate(searchQuery)
	return nil
}

func NewSearchingPage() *searchingPage {
	return &searchingPage{}
}

func (s *searchingPage) searchInAll(query string) (results []source.SearchResult, err error) {
	for _, provider := range providers.Sources {
		r, err := provider.SearchFor(query)
		if err != nil {
			return nil, err
		}
		results = append(results, r...)
	}
	return
}

func (s *searchingPage) newResultList() *tview.List {
	resultList := tview.NewList().
		ShowSecondaryText(false).
		SetSelectedBackgroundColor(tcell.ColorGray)

	// vim-like keybinds (k and j navigation)
	resultList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'k':
			resultList.SetCurrentItem(resultList.GetCurrentItem() - 1)
		case 'j':
			item := resultList.GetCurrentItem() + 1
			if item >= resultList.GetItemCount() {
				item = 0
			}
			resultList.SetCurrentItem(item)
		}
		return event
	})
	return resultList
}

func (s *searchingPage) searchAndThenUpdate(searchQuery string) {
	s.label.SetText(fmt.Sprintf("Searching for %s...", searchQuery))
	results, err := s.searchInAll(searchQuery)
	ui.App.QueueUpdateDraw(func() {
		if err != nil {
			slog.Error("Failed to search", "err", err)
			s.label.SetText("Something went wrong =(")
		}
		s.label.SetText(fmt.Sprintf("Results for %s:", searchQuery))
		for i, result := range results {
			// TODO: is it necessary?
			// limit results to 10
			if i == 10 {
				break
			}
			shortcut := strconv.Itoa(i + 1)
			var details string
			if result.IsLive {
				details = core.FmtEscaping("%s from %s - LIVE", result.Title, result.Artist)
			} else {
				details = core.FmtEscaping("%s from %s - %s", result.Title, result.Artist, result.Length)
			}

			currentResult := result

			s.resultList.AddItem(
				details, "", rune(shortcut[len(shortcut)-1]), func() {
					ui.SwitchPage(ui.PlayingPageName, currentResult)
				},
			)
		}
		s.resultList.AddItem("Cancel", "Press c to cancel", 'c', func() {
			ui.SwitchPage(ui.HomePageName)
		})
	})
}
