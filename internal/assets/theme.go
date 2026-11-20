package assets

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

var (
	hasDarkBG = true

	ColorAccent     color.Color = lipgloss.Color("#ff4b0c")
	ColorForeground color.Color = lipgloss.ANSIColor(15)
	ColorMuted      color.Color = lipgloss.ANSIColor(8)
	ColorBackground color.Color = lipgloss.ANSIColor(0)
)

func SetDarkBackground(dark bool) {
	hasDarkBG = dark
	ld := lipgloss.LightDark(dark)

	ColorForeground = ld(lipgloss.ANSIColor(0), lipgloss.ANSIColor(15))
	ColorMuted = ld(lipgloss.ANSIColor(8), lipgloss.ANSIColor(8))
}

func SetBackground(c color.Color) {
	ColorBackground = c
}
