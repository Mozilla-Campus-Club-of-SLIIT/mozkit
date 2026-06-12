package components

import (
	"fmt"
	"os/exec"
	"strings"

	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/Mozilla-Campus-Club-of-SLIIT/mozkit/internal/assets"
	"github.com/Mozilla-Campus-Club-of-SLIIT/mozkit/internal/engine"
)

var (
	cmdStyle   = lipgloss.NewStyle().Foreground(assets.ColorOrange).Bold(true)
	outStyle   = lipgloss.NewStyle().Foreground(assets.ColorGray)
	errStyle   = lipgloss.NewStyle().Foreground(assets.ColorOrange)
)

type bashOutputMsg struct {
	Index  int
	Output string
	Err    error
}

type ScriptPage struct {
	Model       viewport.Model
	Script      engine.Script
	EnterCount  int
	Confirmed   bool
	Running     bool
	Executed    bool
	actionIndex int
	outputBuf   string
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
		k := msg.String()
		if (p.Running || p.Executed) && (k == "q" || k == "esc") {
			return p, nil
		}
		if !p.Confirmed && !p.Running {
			if k == "enter" {
				p.EnterCount++
				if p.EnterCount >= 3 {
					p.Confirmed = true
				}
			}
		}

	case bashOutputMsg:
		cmdLine := strings.Join(p.Script.Actions[msg.Index].Arguments, " ")
		p.outputBuf += cmdStyle.Render("$ "+cmdLine) + "\n"
		if msg.Err != nil {
			p.outputBuf += errStyle.Render(fmt.Sprintf("(error: %v)", msg.Err)) + "\n"
		}
		if len(msg.Output) > 0 {
			p.outputBuf += outStyle.Render(msg.Output)
			if msg.Output[len(msg.Output)-1] != '\n' {
				p.outputBuf += "\n"
			}
		}
		if msg.Index < len(p.Script.Actions)-1 {
			p.outputBuf += "\n"
		}

		p.Model.SetContent(p.outputBuf)

		next := msg.Index + 1
		if next < len(p.Script.Actions) {
			p.actionIndex = next
			return p, p.runAction(next)
		}
		p.Executed = true
		p.Running = false
	}

	if p.Confirmed && !p.Running && !p.Executed {
		p.Running = true
		p.actionIndex = 0
		p.Model.LeftGutterFunc = func(info viewport.GutterContext) string {
			if info.Soft {
				return "     │ "
			}
			if info.Index >= info.TotalLines {
				return "   ~ │ "
			}
			return fmt.Sprintf("%4d │ ", info.Index+1)
		}
		return p, p.runAction(0)
	}

	p.Model, cmd = p.Model.Update(msg)
	return p, cmd
}

func (p ScriptPage) runAction(index int) tea.Cmd {
	action := p.Script.Actions[index]
	return func() tea.Msg {
		cmd := exec.Command("bash", "-c", strings.Join(action.Arguments, " "))
		output, err := cmd.CombinedOutput()
		return bashOutputMsg{
			Index:  index,
			Output: string(output),
			Err:    err,
		}
	}
}

func (p ScriptPage) View() string {
	if p.Running || p.Executed {
		return p.Model.View()
	}

	var content string

	if p.Confirmed {
		content = "running..."
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
		for i, action := range p.Script.Actions {
			content += fmt.Sprintf("\t%d. %s\n", i+1, action.Description)
		}
	}

	p.Model.SetContent(content)
	return p.Model.View()
}

func NewScript(filePath string, width, height int) ScriptPage {

	script := engine.LoadScript(filePath)

	vp := viewport.New(viewport.WithWidth(width), viewport.WithHeight(height))

	return ScriptPage{
		Model:  vp,
		Script: script,
		EnterCount: -1,
		Confirmed:  false,
	}
}
