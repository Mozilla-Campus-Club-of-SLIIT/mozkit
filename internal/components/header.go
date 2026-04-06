package components

import (
	"strings"

	"github.com/Mozilla-Campus-Club-of-SLIIT/mozkit/internal/assets"
	"github.com/charmbracelet/lipgloss"
)

const decorator = "/"

func leftDecorate(logo string) string {
	const numberOfDecoratorsPerLine = 4
	height := lipgloss.Height(logo)
	lines := make([]string, height)
	for i := range lines {
		lines[i] = strings.Repeat(decorator, numberOfDecoratorsPerLine)
	}
	return lipgloss.NewStyle().Foreground(assets.ColorGray).Bold(true).Render(strings.Join(lines, "\n"))
}

func rightDecorate(leftDecoration, logo string, width int) string {
	// width (Terminal) - 2 (Global Left/Right Padding)
	remainingWidth := width - lipgloss.Width(leftDecoration) - lipgloss.Width(logo) - 4
	if remainingWidth <= 0 {
		return ""
	}
	height := lipgloss.Height(logo)
	lines := make([]string, height)
	for i := range lines {
		lines[i] = strings.Repeat(decorator, remainingWidth)
	}
	return lipgloss.NewStyle().Foreground(assets.ColorGray).Bold(true).Render(strings.Join(lines, "\n"))
}

func Header(width int) string {
	const margin = 2
	logo := assets.RenderLogo()
	leftDecoration := lipgloss.NewStyle().MarginRight(margin).Render(
		leftDecorate(logo))
	rightDecoration := lipgloss.NewStyle().MarginLeft(margin).Render(rightDecorate(leftDecoration, logo, width))
	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		leftDecoration,
		logo,
		rightDecoration,
	)
}
