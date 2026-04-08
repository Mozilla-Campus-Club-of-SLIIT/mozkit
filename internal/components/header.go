package components

import (
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/Mozilla-Campus-Club-of-SLIIT/mozkit/internal/assets"
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

func breadcrumbs(location ...string) string {
	content := "mozkit://"
	if location != nil {
		content += strings.Join(location, "/")
	}
	return lipgloss.NewStyle().Padding(0, 1).Background(assets.ColorOrange).Foreground(assets.ColorWhite).Bold(true).Render(content)
}

func Header(width int, location ...string) string {
	const margin = 2
	logo := assets.RenderLogo()
	leftDecoration := lipgloss.NewStyle().MarginRight(margin).Render(
		leftDecorate(logo))
	rightDecoration := lipgloss.NewStyle().MarginLeft(margin).Render(rightDecorate(leftDecoration, logo, width))
	breadcrumbs := lipgloss.NewStyle().MarginTop(1).Render(breadcrumbs(location...))

	return lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			leftDecoration,
			logo,
			rightDecoration,
		),
		breadcrumbs,
	)

}
