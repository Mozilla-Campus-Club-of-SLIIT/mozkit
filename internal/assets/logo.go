package assets

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func Logo() string {
	const splitIndex = 50
	mozStyle := lipgloss.NewStyle().Foreground(ColorWhite)
	kitStyle := lipgloss.NewStyle().Foreground(ColorOrange)

	logo := []string{
		"██████  ██████  ███████ ███████  ███████████████  ████     ██████    ████   ████ ████ ████",
		"██████  ██████  ███████ ███████  ███████████████  ████   ██████      ████   ████ ████ ████",
		"█  ████  █████  ████       ████          ███████  ████  █████        ████        ████",
		"██  ████  ████  ████       ████      ███████      ████ ████          ████        ████",
		"███  ███  ████  ████       ████   ██████          ████  █████        ████        ████",
		"████  █   ████  ███████ ███████  ███████████████  ████   ██████      ████        ████",
		"████      ████  ███████ ███████  ███████████████  ████     ██████    ████        ████",
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
	logo := Logo()
	return lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.NewStyle().Foreground(ColorOrange).Render("sliitmozilla.org"),
		logo,
	)
}
