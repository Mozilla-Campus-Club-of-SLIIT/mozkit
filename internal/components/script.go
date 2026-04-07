package components

import (
	"fmt"

	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/Mozilla-Campus-Club-of-SLIIT/mozkit/internal/assets"
	"github.com/Mozilla-Campus-Club-of-SLIIT/mozkit/internal/engine"
)

type ScriptPage struct {
	Model      viewport.Model
	Script     engine.Script
	EnterCount int
	Confirmed  bool
}

func (p ScriptPage) SetSize(width, height int) Page {
	p.Model.SetWidth(width)
	p.Model.SetHeight(height)
	return p
}

func (p ScriptPage) Update(msg tea.Msg) (Page, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "enter":
			if !p.Confirmed {
				p.EnterCount++
				if p.EnterCount >= 3 {
					p.Confirmed = true
				}
			}
		}
	}

	p.Model, cmd = p.Model.Update(msg)
	return p, cmd
}

func (p ScriptPage) View() string {
	content := ""

	if p.Confirmed {
		content = "okay, let's do this!"
	} else {
		content = lipgloss.NewStyle().Foreground(assets.ColorWhite).Bold(true).Render(
			"Hey! This script is going to run the following actions. Do you want to continue?\nKeep in mind that there is no undo, and it cannot be stopped in the middle.",
		)

		content += "\n\n"

		content += lipgloss.NewStyle().Foreground(assets.ColorOrange).Bold(true).Render(
			fmt.Sprintf("Press ENTER 3 times to continue (%d/3) | Press ESC to go back", p.EnterCount),
		)

		content += "\n\n"

		content += "Actions:\n"
		var b []byte
		for i, action := range p.Script.Actions {
			b = append(b, "\t"...)
			b = append(b, string(rune(i+1))...)
			b = append(b, ". "...)
			b = append(b, action.Description...)
			b = append(b, "\n"...)
		}
		content += string(b)
	}

	p.Model.SetContent(content)

	return p.Model.View()
}

func NewScript(filePath string, width, height int) ScriptPage {

	script := engine.LoadScript(filePath)

	vp := viewport.New(viewport.WithWidth(width), viewport.WithHeight(height))

	return ScriptPage{
		Model:      vp,
		Script:     script,
		EnterCount: 0,
		Confirmed:  false,
	}
}
