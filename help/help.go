package help

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type KeyMap interface {
	ShortHelp() []key.Binding

	FullHelp() [][]key.Binding
}

type Styles struct {
	Ellipsis lipgloss.Style

	ShortKey       lipgloss.Style
	ShortDesc      lipgloss.Style
	ShortSeparator lipgloss.Style

	FullKey       lipgloss.Style
	FullDesc      lipgloss.Style
	FullSeparator lipgloss.Style
}

func DefaultStyles(isDark bool) Styles { _ = "STUB: not implemented"; return *new(Styles) }

func DefaultDarkStyles() Styles { _ = "STUB: not implemented"; return *new(Styles) }

func DefaultLightStyles() Styles { _ = "STUB: not implemented"; return *new(Styles) }

type Model struct {
	ShowAll bool

	ShortSeparator string
	FullSeparator  string

	Ellipsis string

	Styles Styles

	width int
}

func New() Model { _ = "STUB: not implemented"; return *new(Model) }

func (m Model) Update(_ tea.Msg) (Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(Model), *new(tea.Cmd)
}

func (m Model) View(k KeyMap) string { _ = "STUB: not implemented"; return "" }

func (m *Model) SetWidth(w int) { _ = "STUB: not implemented"; return }

func (m Model) Width() int { _ = "STUB: not implemented"; return 0 }

func (m Model) ShortHelpView(bindings []key.Binding) string { _ = "STUB: not implemented"; return "" }

func (m Model) FullHelpView(groups [][]key.Binding) string { _ = "STUB: not implemented"; return "" }

func (m Model) shouldAddItem(totalWidth, width int) (tail string, ok bool) {
	_ = "STUB: not implemented"
	return "", false
}

func shouldRenderColumn(b []key.Binding) (ok bool) { _ = "STUB: not implemented"; return false }
