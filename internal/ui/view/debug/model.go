package debug

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pauloo27/tuner/internal/core"
	"github.com/pauloo27/tuner/internal/core/logging"
	"github.com/pauloo27/tuner/internal/providers"
	"github.com/pauloo27/tuner/internal/ui/view/root"
)

const headerHeight = 5

type model struct {
	logViewport viewport.Model
}

func NewModel() model {
	logViewport := viewport.New(0, 0)
	return model{logViewport: logViewport}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case root.VisibleMsg:
		m.logViewport.Width, m.logViewport.Height = msg.Width, msg.Height-headerHeight
		m.logViewport.SetContent(logging.MemoryBuffer.String())
		m.logViewport.GotoBottom()
	}

	var viewportCmd tea.Cmd
	m.logViewport, viewportCmd = m.logViewport.Update(msg)

	cmds = append(cmds, viewportCmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	sourcesNames := make([]string, 0, len(providers.Sources))

	for _, source := range providers.Sources {
		sourcesNames = append(sourcesNames, source.Name())
	}

	return fmt.Sprintf(
		"%s\n\n%s\n%s\n%s\n",
		textStyle.Render("Debug"),
		textStyle.Render(fmt.Sprintf(
			"Tuner %s | %s %s %s %s | player %s | sources %v",
			core.Version,
			runtime.Version(), runtime.Compiler, runtime.GOOS, runtime.GOARCH,
			providers.Player.Name(),
			sourcesNames,
		)),
		m.logViewport.View(),
		m.logFooterView(),
	)
}

func (m model) logFooterView() string {
	info := textStyle.Render(fmt.Sprintf("%3.f%%", m.logViewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.logViewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}
