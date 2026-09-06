package viewport

import (
	"cmp"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const (
	defaultHorizontalStep = 6
)

type Option func(*Model)

func WithWidth(w int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithHeight(h int) Option { _ = "STUB: not implemented"; return *new(Option) }

func New(opts ...Option) (m Model) { _ = "STUB: not implemented"; return *new(Model) }

type Model struct {
	width  int
	height int
	KeyMap KeyMap

	SoftWrap bool

	FillHeight bool

	MouseWheelEnabled bool

	MouseWheelDelta int

	yOffset int

	xOffset int

	horizontalStep int

	YPosition int

	Style lipgloss.Style

	LeftGutterFunc GutterFunc

	initialized      bool
	lines            []string
	longestLineWidth int

	HighlightStyle lipgloss.Style

	SelectedHighlightStyle lipgloss.Style

	StyleLineFunc func(int) lipgloss.Style

	highlights []highlightInfo
	hiIdx      int
}

type GutterFunc func(GutterContext) string

var NoGutter = func(GutterContext) string { return "" }

type GutterContext struct {
	Index int

	TotalLines int

	Soft bool
}

func (m *Model) setInitialValues() { _ = "STUB: not implemented"; return }

func (m Model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m Model) Height() int { _ = "STUB: not implemented"; return 0 }

func (m *Model) SetHeight(h int) { _ = "STUB: not implemented"; return }

func (m Model) Width() int { _ = "STUB: not implemented"; return 0 }

func (m *Model) SetWidth(w int) { _ = "STUB: not implemented"; return }

func (m Model) AtTop() bool { _ = "STUB: not implemented"; return false }

func (m Model) AtBottom() bool { _ = "STUB: not implemented"; return false }

func (m Model) PastBottom() bool { _ = "STUB: not implemented"; return false }

func (m Model) ScrollPercent() float64 { _ = "STUB: not implemented"; return 0 }

func (m Model) HorizontalScrollPercent() float64 { _ = "STUB: not implemented"; return 0 }

func (m *Model) SetContent(s string) { _ = "STUB: not implemented"; return }

func (m *Model) SetContentLines(lines []string) { _ = "STUB: not implemented"; return }

func (m Model) GetContent() string { _ = "STUB: not implemented"; return "" }

func (m Model) calculateLine(yoffset int) (total, ridx, voffset int) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

func (m Model) maxYOffset() int { _ = "STUB: not implemented"; return 0 }

func (m Model) maxXOffset() int { _ = "STUB: not implemented"; return 0 }

func (m Model) maxWidth() int { _ = "STUB: not implemented"; return 0 }

func (m Model) maxHeight() int { _ = "STUB: not implemented"; return 0 }

func (m Model) visibleLines() (lines []string) { _ = "STUB: not implemented"; return nil }

func (m Model) styleLines(lines []string, offset int) []string {
	_ = "STUB: not implemented"
	return nil
}

func (m Model) highlightLines(lines []string, offset int) []string {
	_ = "STUB: not implemented"
	return nil
}

func (m Model) softWrap(lines []string, maxWidth, maxHeight, total, ridx, voffset int) []string {
	_ = "STUB: not implemented"
	return nil
}

func (m Model) setupGutter(lines []string, total, ridx int) []string {
	_ = "STUB: not implemented"
	return nil
}

func (m *Model) SetYOffset(n int) { _ = "STUB: not implemented"; return }

func (m *Model) YOffset() int { _ = "STUB: not implemented"; return 0 }

func (m *Model) EnsureVisible(line, colstart, colend int) { _ = "STUB: not implemented"; return }

func (m *Model) PageDown() { _ = "STUB: not implemented"; return }

func (m *Model) PageUp() { _ = "STUB: not implemented"; return }

func (m *Model) HalfPageDown() { _ = "STUB: not implemented"; return }

//nolint:mnd

func (m *Model) HalfPageUp() { _ = "STUB: not implemented"; return }

//nolint:mnd

func (m *Model) ScrollDown(n int) { _ = "STUB: not implemented"; return }

func (m *Model) ScrollUp(n int) { _ = "STUB: not implemented"; return }

func (m *Model) SetHorizontalStep(n int) { _ = "STUB: not implemented"; return }

func (m *Model) XOffset() int { _ = "STUB: not implemented"; return 0 }

func (m *Model) SetXOffset(n int) { _ = "STUB: not implemented"; return }

func (m *Model) ScrollLeft(n int) { _ = "STUB: not implemented"; return }

func (m *Model) ScrollRight(n int) { _ = "STUB: not implemented"; return }

func (m Model) TotalLineCount() int { _ = "STUB: not implemented"; return 0 }

func (m Model) VisibleLineCount() int { _ = "STUB: not implemented"; return 0 }

func (m *Model) GotoTop() (lines []string) { _ = "STUB: not implemented"; return nil }

func (m *Model) GotoBottom() (lines []string) { _ = "STUB: not implemented"; return nil }

func (m *Model) SetHighlights(matches [][]int) { _ = "STUB: not implemented"; return }

func (m *Model) ClearHighlights() { _ = "STUB: not implemented"; return }

func (m *Model) showHighlight() { _ = "STUB: not implemented"; return }

func (m *Model) HighlightNext() { _ = "STUB: not implemented"; return }

func (m *Model) HighlightPrevious() { _ = "STUB: not implemented"; return }

func (m Model) findNearestMatch() int { _ = "STUB: not implemented"; return 0 }

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(Model), *new(tea.Cmd)
}

func (m Model) updateAsModel(msg tea.Msg) Model { _ = "STUB: not implemented"; return *new(Model) }

func (m Model) View() string { _ = "STUB: not implemented"; return "" }

func clamp[T cmp.Ordered](v, low, high T) T { _ = "STUB: not implemented"; return *new(T) }

func maxLineWidth(lines []string) int { _ = "STUB: not implemented"; return 0 }
