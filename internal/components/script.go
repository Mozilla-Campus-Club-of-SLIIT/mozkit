package components

import (
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"github.com/Mozilla-Campus-Club-of-SLIIT/mozkit/internal/engine"
)

type ScriptPage struct {
	Model viewport.Model
}

func (p ScriptPage) SetSize(width, height int) Page {
	p.Model.SetWidth(width)
	p.Model.SetHeight(height)
	return p
}

func (p ScriptPage) Update(msg tea.Msg) (Page, tea.Cmd) {
	var cmd tea.Cmd
	p.Model, cmd = p.Model.Update(msg)
	return p, cmd
}

func (p ScriptPage) View() string {
	return p.Model.View()
}

func NewScript(filePath string, width, height int) ScriptPage {
	script := engine.LoadScript(filePath)
	content := "Title: " + script.Title + "\n\n" + "Description: " + script.Description + "\n\nActions:\n"
	var b []byte
	for i, action := range script.Actions {
		b = append(b, "\t"...)
		b = append(b, string(rune(i+1))...)
		b = append(b, ". "...)
		b = append(b, action.Description...)
		b = append(b, "\n"...)
	}
	content += string(b)

	vp := viewport.New(viewport.WithWidth(width), viewport.WithHeight(height))
	vp.SetContent(content)

	return ScriptPage{
		Model: vp,
	}
}
