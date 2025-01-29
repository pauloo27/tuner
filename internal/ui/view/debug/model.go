package debug

import (
	"fmt"
	"runtime"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pauloo27/tuner/internal/core"
	"github.com/pauloo27/tuner/internal/providers"
	"github.com/pauloo27/tuner/internal/ui/view/root"
)

type model struct {
	counter int
}

func NewModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg.(type) {
	case root.VisibleMsg:
		cmd = tea.Tick(1*time.Second, func(time.Time) tea.Msg { return incMsg{} })
	case incMsg:
		m.counter++
		// tick only does it once. If we need to keep the cmd in a loop, we gotta
		// send it again
		cmd = tea.Tick(1*time.Second, func(time.Time) tea.Msg { return incMsg{} })
	}
	return m, cmd
}

func (m model) View() string {
	sourcesNames := make([]string, 0, len(providers.Sources))

	for _, source := range providers.Sources {
		sourcesNames = append(sourcesNames, source.Name())
	}

	return fmt.Sprintf(
		"%s\n\n%s\n",
		textStyle.Render("Debug"),
		textStyle.Render(fmt.Sprintf(
			"Tuner %s | %s %s %s %s | player %s | sources %v | counter %d",
			core.Version,
			runtime.Version(), runtime.Compiler, runtime.GOOS, runtime.GOARCH,
			providers.Player.Name(),
			sourcesNames,
			m.counter,
		)),
	)
}
