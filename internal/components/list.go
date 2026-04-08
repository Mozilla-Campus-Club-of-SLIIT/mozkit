package components

import (
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"github.com/Mozilla-Campus-Club-of-SLIIT/mozkit/internal/assets"
)

type ListPage struct {
	Model list.Model
}

func (p ListPage) SetSize(width, height int) Page {
	p.Model.SetSize(width, height)
	return p
}

func (p ListPage) Update(msg tea.Msg) (Page, tea.Cmd) {
	var cmd tea.Cmd
	p.Model, cmd = p.Model.Update(msg)
	return p, cmd
}

func (p ListPage) View() string {
	return p.Model.View()
}

func NewList(items []list.Item, width, height int) list.Model {
	delegate := list.NewDefaultDelegate()
	delegate.Styles.NormalTitle = delegate.Styles.NormalTitle.
		Foreground(assets.ColorGray).
		Bold(false)
	delegate.Styles.NormalDesc = delegate.Styles.NormalDesc.
		Foreground(assets.ColorGray).
		Bold(false)

	delegate.Styles.SelectedTitle = delegate.Styles.SelectedTitle.
		Foreground(assets.ColorOrange).
		Bold(true).
		BorderForeground(assets.ColorOrange)

	delegate.Styles.SelectedDesc = delegate.Styles.SelectedDesc.
		Foreground(assets.ColorWhite).
		Bold(true).
		BorderForeground(assets.ColorOrange)

	l := list.New(items, delegate, width, height)

	l.SetShowTitle(false)
	l.SetShowStatusBar(false)
	l.SetShowHelp(false)
	l.KeyMap.Quit.Unbind()
	l.KeyMap.ForceQuit.Unbind()
	// l.SetFilteringEnabled(false)

	return l
}
