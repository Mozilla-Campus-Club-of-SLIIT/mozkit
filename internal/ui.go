package internal

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/Mozilla-Campus-Club-of-SLIIT/mozkit/internal/components"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	width int
}

func NewModel() *Model {
	return &Model{}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width

	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *Model) View() tea.View {
	header := components.Header()
	tmp := fmt.Sprintf("Termianl Width is %d", m.width)
	view := tea.NewView(
		lipgloss.JoinVertical(
			lipgloss.Top,
			header,
			tmp,
		),
	)
	view.AltScreen = true
	return view
}
