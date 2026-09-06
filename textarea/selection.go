package textarea

type Position struct {
	Row int
	Col int
}

func (p Position) before(q Position) bool { _ = "STUB: not implemented"; return false }

func (m Model) gutterWidth() int { _ = "STUB: not implemented"; return 0 }

func runeIndexForColumn(runes []rune, col int) int { _ = "STUB: not implemented"; return 0 }

func (m Model) PositionAt(x, y int) Position { _ = "STUB: not implemented"; return *new(Position) }

func (m *Model) BeginSelection(x, y int) { _ = "STUB: not implemented"; return }

func (m *Model) ExtendSelection(x, y int) { _ = "STUB: not implemented"; return }

func (m *Model) EndSelection() { _ = "STUB: not implemented"; return }

func (m *Model) SelectAll() { _ = "STUB: not implemented"; return }

func (m *Model) ClearSelection() { _ = "STUB: not implemented"; return }

func (m Model) HasSelection() bool { _ = "STUB: not implemented"; return false }

func (m Model) Selection() (start, end Position, ok bool) {
	_ = "STUB: not implemented"
	return *new(Position), *new(Position), false
}

func (m Model) SelectedText() string { _ = "STUB: not implemented"; return "" }

func (m *Model) selectFrom(anchor, head Position) { _ = "STUB: not implemented"; return }

func (m *Model) moveCursorTo(pos Position) { _ = "STUB: not implemented"; return }

func (m Model) selectionSpanFor(row, base, length int) (from, to int, ok bool) {
	_ = "STUB: not implemented"
	return 0, 0, false
}
