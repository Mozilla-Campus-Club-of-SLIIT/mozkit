package components

import (
	"strings"

	"github.com/Mozilla-Campus-Club-of-SLIIT/mozkit/internal/assets"
	"github.com/charmbracelet/lipgloss"
)

var KeyBindings = [][2]string{
	{"↑/k", "up"},
	{"↓/j", "down"},
	{"enter", "select"},
	{"esc", "back"},
	{"q", "quit"},
}

func Footer() string {
	var parts []string

	keyStyle := lipgloss.NewStyle().Foreground(assets.ColorGray).Bold(true)
	actionStyle := lipgloss.NewStyle().Foreground(assets.ColorGray)

	for _, binding := range KeyBindings {
		keycap := keyStyle.Render(binding[0])
		action := actionStyle.Render(": " + binding[1])
		parts = append(parts, keycap+action)
	}
	return strings.Join(parts, "  ")
}
