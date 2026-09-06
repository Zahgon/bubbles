package list

import (
	"io"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type DefaultItemStyles struct {
	NormalTitle lipgloss.Style
	NormalDesc  lipgloss.Style

	SelectedTitle lipgloss.Style
	SelectedDesc  lipgloss.Style

	DimmedTitle lipgloss.Style
	DimmedDesc  lipgloss.Style

	FilterMatch lipgloss.Style
}

func NewDefaultItemStyles(isDark bool) (s DefaultItemStyles) {
	_ = "STUB: not implemented"
	return *new(DefaultItemStyles)
}

//nolint:mnd

//nolint:mnd

type DefaultItem interface {
	Item
	Title() string
	Description() string
}

type DefaultDelegate struct {
	ShowDescription bool
	Styles          DefaultItemStyles
	UpdateFunc      func(tea.Msg, *Model) tea.Cmd
	ShortHelpFunc   func() []key.Binding
	FullHelpFunc    func() [][]key.Binding
	height          int
	spacing         int
}

func NewDefaultDelegate() DefaultDelegate { _ = "STUB: not implemented"; return *new(DefaultDelegate) }

func (d *DefaultDelegate) SetHeight(i int) { _ = "STUB: not implemented"; return }

func (d DefaultDelegate) Height() int { _ = "STUB: not implemented"; return 0 }

func (d *DefaultDelegate) SetSpacing(i int) { _ = "STUB: not implemented"; return }

func (d DefaultDelegate) Spacing() int { _ = "STUB: not implemented"; return 0 }

func (d DefaultDelegate) Update(msg tea.Msg, m *Model) tea.Cmd {
	_ = "STUB: not implemented"
	return *new(tea.Cmd)
}

func (d DefaultDelegate) Render(w io.Writer, m Model, index int, item Item) {
	_ = "STUB: not implemented"
	return
}

//nolint: errcheck

//nolint: errcheck

func (d DefaultDelegate) ShortHelp() []key.Binding { _ = "STUB: not implemented"; return nil }

func (d DefaultDelegate) FullHelp() [][]key.Binding { _ = "STUB: not implemented"; return nil }
