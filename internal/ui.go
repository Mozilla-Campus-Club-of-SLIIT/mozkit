package internal

import (
	"fmt"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/Mozilla-Campus-Club-of-SLIIT/mozkit/internal/assets"
	"github.com/Mozilla-Campus-Club-of-SLIIT/mozkit/internal/components"
	"github.com/Mozilla-Campus-Club-of-SLIIT/mozkit/internal/engine"
)

type Model struct {
	width   int
	height  int
	content components.Page
	stack   []engine.Item
}

func NewModel() *Model {
	engineItems := engine.CollectionList()
	listItems := make([]list.Item, len(engineItems))
	for i, it := range engineItems {
		listItems[i] = it
	}
	menu := components.NewList(listItems, 0, 0)

	return &Model{
		content: components.ListPage{Model: menu},
	}
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
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			if page, ok := m.content.(components.ListPage); ok {
				if item, ok := page.Model.SelectedItem().(engine.Item); ok {
					if len(m.stack) == 0 {
						m.stack = append(m.stack, item)

						newItems := engine.ScriptList(item.Target)
						listItems := make([]list.Item, len(newItems))
						for i, it := range newItems {
							listItems[i] = it
						}

						cmd := page.Model.SetItems(listItems)
						page.Model.ResetSelected()
						m.content = page
						return m, cmd
					} else {
						m.stack = append(m.stack, item)
						m.content = components.NewScript(item.Target, m.width, m.height)
					}
				}
			}

		case "esc":
			if page, ok := m.content.(components.ListPage); ok {
				if page.Model.FilterState() == list.Filtering {
					break
				}

				if len(m.stack) > 0 {
					m.stack = m.stack[:len(m.stack)-1]

					rootItems := engine.CollectionList()
					listItems := make([]list.Item, len(rootItems))
					for i, it := range rootItems {
						listItems[i] = it
					}

					cmd := page.Model.SetItems(listItems)
					page.Model.ResetSelected()
					m.content = page
					return m, cmd
				} else if len(m.stack) == 0 {
					return m, nil
				}
			} else if _, ok := m.content.(components.ScriptPage); ok {
				m.stack = m.stack[:len(m.stack)-1]

				var listItems []list.Item
				if len(m.stack) > 0 {
					newItems := engine.ScriptList(m.stack[len(m.stack)-1].Target)
					listItems = make([]list.Item, len(newItems))
					for i, it := range newItems {
						listItems[i] = it
					}
				} else {
					rootItems := engine.CollectionList()
					listItems = make([]list.Item, len(rootItems))
					for i, it := range rootItems {
						listItems[i] = it
					}
				}
				menu := components.NewList(listItems, 0, 0)
				m.content = components.ListPage{Model: menu}
			}
		}

	}
	var cmd tea.Cmd
	m.content, cmd = m.content.Update(msg)
	return m, cmd
}

func (m *Model) View() tea.View {
	content := sizeCheck(m.width, m.height, func() string {

		var location []string
		for _, item := range m.stack {
			location = append(location, item.TitleStr)
		}
		header := components.Header(m.width, location...)
		footer := components.Footer()

		// m.height (total)
		// - 2 for the outer Padding(1) on top and bottom
		// - Height of Header
		// - Height of Footer
		// - 2 for the two "" strings we are adding to the layout!
		spaceBetweenUs := m.height - 4 - lipgloss.Height(header) - lipgloss.Height(footer)

		m.content = m.content.SetSize(m.width, spaceBetweenUs)

		content := lipgloss.JoinVertical(
			lipgloss.Left,
			header,
			"",
			m.content.View(),
			"",
			footer,
		)

		return lipgloss.NewStyle().Padding(1).Render(content)
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
