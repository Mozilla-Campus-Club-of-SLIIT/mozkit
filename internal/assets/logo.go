package assets

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
)

func logo() string {
	const splitIndex = 50
	mozStyle := lipgloss.NewStyle().Foreground(ColorForeground)
	kitStyle := lipgloss.NewStyle().Foreground(ColorAccent)

	logo := []string{
		"██████  ██████  ███████ ███████  ███████████████  ████     ██████  ████  ████ ████ ████",
		"██████  ██████  ███████ ███████  ███████████████  ████   ██████    ████  ████ ████ ████",
		"█  ████  █████  ████       ████          ███████  ████  █████      ████       ████",
		"██  ████  ████  ████       ████      ███████      ████ ████        ████       ████",
		"███  ███  ████  ████       ████   ██████          ████  █████      ████       ████",
		"████  █   ████  ███████ ███████  ███████████████  ████   ██████    ████       ████",
		"████      ████  ███████ ███████  ███████████████  ████     ██████  ████       ████",
	}

	for i, line := range logo {
		runes := []rune(line)
		mozPart := mozStyle.Render(string(runes[:splitIndex]))
		kitPart := kitStyle.Render(string(runes[splitIndex:]))

		logo[i] = mozPart + kitPart
	}

	return fmt.Sprintf("%s", strings.Join(logo, "\n"))
}
func RenderLogo() string {
	logo := logo()
	return lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.NewStyle().Foreground(ColorAccent).Bold(true).Render("sliitmozilla.org"),
		logo,
	)
}
