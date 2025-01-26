package debug

import (
	"fmt"
	"runtime"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pauloo27/tuner/internal/core"
	"github.com/pauloo27/tuner/internal/providers"
)

type model struct {
}

func NewModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
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
			"Tuner %s | %s %s %s %s | player %s | sources %v",
			core.Version,
			runtime.Version(), runtime.Compiler, runtime.GOOS, runtime.GOARCH,
			providers.Player.Name(),
			sourcesNames,
		)),
	)
}
