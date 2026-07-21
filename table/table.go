package table

import (
	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type Model struct {
	KeyMap KeyMap
	Help   help.Model

	cols   []Column
	rows   []Row
	cursor int
	focus  bool
	styles Styles

	viewport viewport.Model
	start    int
	end      int
}

type Row []string

type Column struct {
	Title string
	Width int
}

type KeyMap struct {
	LineUp       key.Binding
	LineDown     key.Binding
	PageUp       key.Binding
	PageDown     key.Binding
	HalfPageUp   key.Binding
	HalfPageDown key.Binding
	GotoTop      key.Binding
	GotoBottom   key.Binding
}

func (km KeyMap) ShortHelp() []key.Binding { _ = "STUB: not implemented"; return nil }

func (km KeyMap) FullHelp() [][]key.Binding { _ = "STUB: not implemented"; return nil }

func DefaultKeyMap() KeyMap { _ = "STUB: not implemented"; return *new(KeyMap) }

type Styles struct {
	Header   lipgloss.Style
	Cell     lipgloss.Style
	Selected lipgloss.Style
}

func DefaultStyles() Styles { _ = "STUB: not implemented"; return *new(Styles) }

func (m *Model) SetStyles(s Styles) { _ = "STUB: not implemented"; return }

type Option func(*Model)

func New(opts ...Option) Model { _ = "STUB: not implemented"; return *new(Model) }

//nolint:mnd

func WithColumns(cols []Column) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithRows(rows []Row) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithHeight(h int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithWidth(w int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithFocused(f bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithStyles(s Styles) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithKeyMap(km KeyMap) Option { _ = "STUB: not implemented"; return *new(Option) }

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(Model), *new(tea.Cmd)
}

//nolint:mnd

//nolint:mnd

func (m Model) Focused() bool { _ = "STUB: not implemented"; return false }

func (m *Model) Focus() { _ = "STUB: not implemented"; return }

func (m *Model) Blur() { _ = "STUB: not implemented"; return }

func (m Model) View() string { _ = "STUB: not implemented"; return "" }

func (m Model) HelpView() string { _ = "STUB: not implemented"; return "" }

func (m *Model) UpdateViewport() { _ = "STUB: not implemented"; return }

func (m Model) SelectedRow() Row { _ = "STUB: not implemented"; return *new(Row) }

func (m Model) Rows() []Row { _ = "STUB: not implemented"; return nil }

func (m Model) Columns() []Column { _ = "STUB: not implemented"; return nil }

func (m *Model) SetRows(r []Row) { _ = "STUB: not implemented"; return }

func (m *Model) SetColumns(c []Column) { _ = "STUB: not implemented"; return }

func (m *Model) SetWidth(w int) { _ = "STUB: not implemented"; return }

func (m *Model) SetHeight(h int) { _ = "STUB: not implemented"; return }

func (m Model) Height() int { _ = "STUB: not implemented"; return 0 }

func (m Model) Width() int { _ = "STUB: not implemented"; return 0 }

func (m Model) Cursor() int { _ = "STUB: not implemented"; return 0 }

func (m *Model) SetCursor(n int) { _ = "STUB: not implemented"; return }

func (m *Model) MoveUp(n int) { _ = "STUB: not implemented"; return }

func (m *Model) MoveDown(n int) { _ = "STUB: not implemented"; return }

func (m *Model) GotoTop() { _ = "STUB: not implemented"; return }

func (m *Model) GotoBottom() { _ = "STUB: not implemented"; return }

func (m *Model) FromValues(value, separator string) {
	_ = "STUB: not implemented"
	//nolint:prealloc
	return
}

func (m Model) headersView() string { _ = "STUB: not implemented"; return "" }

func (m *Model) renderRow(r int) string { _ = "STUB: not implemented"; return "" }

func clamp(v, low, high int) int { _ = "STUB: not implemented"; return 0 }
