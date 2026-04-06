package internal

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/Mozilla-Campus-Club-of-SLIIT/mozkit/internal/assets"
	"github.com/Mozilla-Campus-Club-of-SLIIT/mozkit/internal/components"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	width  int
	height int
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
		m.height = msg.Height

	case tea.KeyPressMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *Model) View() tea.View {
	content := sizeCheck(m.width, m.height, func() string {

		topSection := lipgloss.JoinVertical(
			lipgloss.Left,
			components.Header(m.width),
			"",
			fmt.Sprintf("Terminal Width is %d, Height is %d", m.width, m.height),
		)

		footer := components.Footer()

		//? Subtract 2 from height to account for Padding(1) on top and bottom
		spaceBetweenUs := m.height - 2 - lipgloss.Height(topSection) - lipgloss.Height(footer)

		normalContent := lipgloss.JoinVertical(
			lipgloss.Top,
			topSection,
			lipgloss.NewStyle().Height(spaceBetweenUs).Render(""),
			footer,
		)

		return lipgloss.NewStyle().Padding(1).Render(normalContent)
	})

	view := tea.NewView(content)
	view.AltScreen = true
	return view
}

func sizeCheck(width, height int, content func() string) string {
	const minWidth = 101
	const minHeight = 21

	if width <= minWidth-1 || height <= minHeight-1 {

		title := lipgloss.NewStyle().Foreground(assets.ColorWhite).Bold(true).Render("I needs a little more room to stretch my legs!")

		stats := lipgloss.NewStyle().Foreground(assets.ColorGray).Align(lipgloss.Center).Render(
			fmt.Sprintf("What you gave me: %dx%d\nWhat I need(at least): %dx%d", width, height, minWidth, minHeight),
		)

		instruction := lipgloss.NewStyle().Foreground(assets.ColorOrange).Bold(true).Render("Could you please resize your terminal or decrease your font size?")

		contentBox := lipgloss.JoinVertical(
			lipgloss.Center,
			title,
			"",
			stats,
			"",
			instruction,
		)

		return lipgloss.NewStyle().
			Width(width).
			Height(height).
			Align(lipgloss.Center, lipgloss.Center).
			Render(contentBox)

	}

	return content()
}
