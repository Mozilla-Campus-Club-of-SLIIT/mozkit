package components

import tea "charm.land/bubbletea/v2"

type Page interface {
	Update(msg tea.Msg) (Page, tea.Cmd)
	View() string
	SetSize(width, height int) Page
}
