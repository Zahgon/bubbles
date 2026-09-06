package textarea

import (
	"image/color"
	"time"

	"charm.land/bubbles/v2/cursor"
	"charm.land/bubbles/v2/internal/memoization"
	"charm.land/bubbles/v2/internal/runeutil"
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const (
	minHeight        = 1
	defaultHeight    = 6
	defaultWidth     = 40
	defaultCharLimit = 0
	defaultMaxHeight = 99
	defaultMaxWidth  = 500

	maxLines = 10000
)

type (
	pasteMsg    string
	pasteErrMsg struct{ error }
	copyMsg     string
	copyErrMsg  struct{ error }
)

type KeyMap struct {
	CharacterBackward       key.Binding
	CharacterForward        key.Binding
	DeleteAfterCursor       key.Binding
	DeleteBeforeCursor      key.Binding
	DeleteCharacterBackward key.Binding
	DeleteCharacterForward  key.Binding
	DeleteWordBackward      key.Binding
	DeleteWordForward       key.Binding
	InsertNewline           key.Binding
	LineEnd                 key.Binding
	LineNext                key.Binding
	LinePrevious            key.Binding
	LineStart               key.Binding
	PageUp                  key.Binding
	PageDown                key.Binding
	Paste                   key.Binding
	WordBackward            key.Binding
	WordForward             key.Binding
	InputBegin              key.Binding
	InputEnd                key.Binding

	UppercaseWordForward  key.Binding
	LowercaseWordForward  key.Binding
	CapitalizeWordForward key.Binding

	TransposeCharacterBackward key.Binding

	SelectCharacterForward  key.Binding
	SelectCharacterBackward key.Binding
	SelectWordForward       key.Binding
	SelectWordBackward      key.Binding
	SelectLineUp            key.Binding
	SelectLineDown          key.Binding
	SelectAll               key.Binding
	CopySelection           key.Binding
}

func DefaultKeyMap() KeyMap { _ = "STUB: not implemented"; return *new(KeyMap) }

type LineInfo struct {
	Width int

	CharWidth int

	Height int

	StartColumn int

	ColumnOffset int

	RowOffset int

	CharOffset int
}

type PromptInfo struct {
	LineNumber int
	Focused    bool
}

type CursorStyle struct {
	Color color.Color

	Shape tea.CursorShape

	Blink bool

	BlinkSpeed time.Duration
}

type Styles struct {
	Focused StyleState
	Blurred StyleState
	Cursor  CursorStyle
}

type StyleState struct {
	Base             lipgloss.Style
	Text             lipgloss.Style
	LineNumber       lipgloss.Style
	CursorLineNumber lipgloss.Style
	CursorLine       lipgloss.Style
	EndOfBuffer      lipgloss.Style
	Placeholder      lipgloss.Style
	Prompt           lipgloss.Style

	Selection lipgloss.Style
}

func (s StyleState) computedCursorLine() lipgloss.Style {
	_ = "STUB: not implemented"
	return *new(lipgloss.Style)
}

func (s StyleState) computedCursorLineNumber() lipgloss.Style {
	_ = "STUB: not implemented"
	return *new(lipgloss.Style)
}

func (s StyleState) computedEndOfBuffer() lipgloss.Style {
	_ = "STUB: not implemented"
	return *new(lipgloss.Style)
}

func (s StyleState) computedLineNumber() lipgloss.Style {
	_ = "STUB: not implemented"
	return *new(lipgloss.Style)
}

func (s StyleState) computedPlaceholder() lipgloss.Style {
	_ = "STUB: not implemented"
	return *new(lipgloss.Style)
}

func (s StyleState) computedPrompt() lipgloss.Style {
	_ = "STUB: not implemented"
	return *new(lipgloss.Style)
}

func (s StyleState) computedText() lipgloss.Style {
	_ = "STUB: not implemented"
	return *new(lipgloss.Style)
}

func (s StyleState) computedSelection() lipgloss.Style {
	_ = "STUB: not implemented"
	return *new(lipgloss.Style)
}

type line struct {
	runes []rune
	width int
}

func (w line) Hash() string { _ = "STUB: not implemented"; return "" }

type Model struct {
	Err error

	cache *memoization.MemoCache[line, [][]rune]

	Prompt string

	Placeholder string

	ShowLineNumbers bool

	EndOfBufferCharacter rune

	KeyMap KeyMap

	virtualCursor cursor.Model

	CharLimit int

	MaxHeight int

	MaxWidth int

	DynamicHeight bool

	MinHeight int

	MaxContentHeight int

	styles Styles

	useVirtualCursor bool

	promptFunc func(PromptInfo) string

	promptWidth int

	width int

	height int

	value [][]rune

	focus bool

	col int

	row int

	lastCharOffset int

	viewport *viewport.Model

	rsan runeutil.Sanitizer

	selAnchor Position
	selHead   Position

	hasSelection bool

	selecting bool
}

func New() Model { _ = "STUB: not implemented"; return *new(Model) }

func DefaultStyles(isDark bool) Styles { _ = "STUB: not implemented"; return *new(Styles) }

func DefaultLightStyles() Styles { _ = "STUB: not implemented"; return *new(Styles) }

func DefaultDarkStyles() Styles { _ = "STUB: not implemented"; return *new(Styles) }

func (m Model) Styles() Styles { _ = "STUB: not implemented"; return *new(Styles) }

func (m *Model) SetStyles(s Styles) { _ = "STUB: not implemented"; return }

func (m Model) VirtualCursor() bool { _ = "STUB: not implemented"; return false }

func (m *Model) SetVirtualCursor(v bool) { _ = "STUB: not implemented"; return }

func (m *Model) updateVirtualCursorStyle() { _ = "STUB: not implemented"; return }

func (m *Model) SetValue(s string) { _ = "STUB: not implemented"; return }

func (m *Model) InsertString(s string) { _ = "STUB: not implemented"; return }

func (m *Model) InsertRune(r rune) { _ = "STUB: not implemented"; return }

func (m *Model) insertRunesFromUserInput(runes []rune) { _ = "STUB: not implemented"; return }

func (m Model) Value() string { _ = "STUB: not implemented"; return "" }

func (m *Model) Length() int { _ = "STUB: not implemented"; return 0 }

func (m *Model) LineCount() int { _ = "STUB: not implemented"; return 0 }

func (m Model) Line() int { _ = "STUB: not implemented"; return 0 }

func (m Model) Column() int { _ = "STUB: not implemented"; return 0 }

func (m Model) ScrollYOffset() int { _ = "STUB: not implemented"; return 0 }

func (m Model) ScrollPercent() float64 { _ = "STUB: not implemented"; return 0 }

func (m *Model) setCursorLineRelative(delta int) { _ = "STUB: not implemented"; return }

//nolint:nestif

func (m *Model) CursorDown() { _ = "STUB: not implemented"; return }

func (m *Model) CursorUp() { _ = "STUB: not implemented"; return }

func (m *Model) SetCursorColumn(col int) { _ = "STUB: not implemented"; return }

func (m *Model) CursorStart() { _ = "STUB: not implemented"; return }

func (m *Model) CursorEnd() { _ = "STUB: not implemented"; return }

func (m Model) Focused() bool { _ = "STUB: not implemented"; return false }

func (m Model) activeStyle() *StyleState { _ = "STUB: not implemented"; return nil }

func (m *Model) Focus() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m *Model) Blur() { _ = "STUB: not implemented"; return }

func (m *Model) Reset() { _ = "STUB: not implemented"; return }

func (m *Model) Word() string { _ = "STUB: not implemented"; return "" }

func (m *Model) san() runeutil.Sanitizer {
	_ = "STUB: not implemented"
	return *new(runeutil.Sanitizer)
}

func (m *Model) deleteBeforeCursor() { _ = "STUB: not implemented"; return }

func (m *Model) deleteAfterCursor() { _ = "STUB: not implemented"; return }

func (m *Model) transposeLeft() { _ = "STUB: not implemented"; return }

func (m *Model) deleteWordLeft() { _ = "STUB: not implemented"; return }

func (m *Model) deleteWordRight() { _ = "STUB: not implemented"; return }

func (m *Model) characterRight() { _ = "STUB: not implemented"; return }

func (m *Model) characterLeft(insideLine bool) { _ = "STUB: not implemented"; return }

func (m *Model) wordLeft() { _ = "STUB: not implemented"; return }

func (m *Model) wordRight() { _ = "STUB: not implemented"; return }

func (m *Model) doWordRight(fn func(charIdx int, pos int)) { _ = "STUB: not implemented"; return }

func (m *Model) uppercaseRight() { _ = "STUB: not implemented"; return }

func (m *Model) lowercaseRight() { _ = "STUB: not implemented"; return }

func (m *Model) capitalizeRight() { _ = "STUB: not implemented"; return }

func (m Model) LineInfo() LineInfo { _ = "STUB: not implemented"; return *new(LineInfo) }

func (m *Model) repositionView() { _ = "STUB: not implemented"; return }

func (m Model) Width() int { _ = "STUB: not implemented"; return 0 }

func (m *Model) MoveToBegin() { _ = "STUB: not implemented"; return }

func (m *Model) MoveToEnd() { _ = "STUB: not implemented"; return }

func (m *Model) PageUp() { _ = "STUB: not implemented"; return }

func (m *Model) PageDown() { _ = "STUB: not implemented"; return }

func (m *Model) SetWidth(w int) { _ = "STUB: not implemented"; return }

func (m *Model) SetPromptFunc(promptWidth int, fn func(PromptInfo) string) {
	_ = "STUB: not implemented"
	return
}

func (m Model) Height() int { _ = "STUB: not implemented"; return 0 }

func (m *Model) SetHeight(h int) { _ = "STUB: not implemented"; return }

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(Model), *new(tea.Cmd)
}

func (m *Model) view() string { _ = "STUB: not implemented"; return "" }

func (m Model) View() string { _ = "STUB: not implemented"; return "" }

func (m Model) promptView(displayLine int) (prompt string) { _ = "STUB: not implemented"; return "" }

func (m Model) lineNumberView(n int, isCursorLine bool) (str string) {
	_ = "STUB: not implemented"
	return ""
}

func (m Model) placeholderView() string { _ = "STUB: not implemented"; return "" }

func Blink() tea.Msg { _ = "STUB: not implemented"; return *new(tea.Msg) }

func (m Model) Cursor() *tea.Cursor { _ = "STUB: not implemented"; return nil }

func (m *Model) deleteSelection() { _ = "STUB: not implemented"; return }

func (m *Model) DeleteSelection() { _ = "STUB: not implemented"; return }

func (m *Model) CopySelection() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m *Model) startKeyboardSelection() { _ = "STUB: not implemented"; return }

func (m *Model) updateKeyboardSelection() { _ = "STUB: not implemented"; return }

func (m Model) memoizedWrap(runes []rune, width int) [][]rune {
	_ = "STUB: not implemented"
	return nil
}

func (m Model) cursorLineNumber() int { _ = "STUB: not implemented"; return 0 }

func (m *Model) totalVisualLines() int { _ = "STUB: not implemented"; return 0 }

func (m *Model) recalculateHeight() { _ = "STUB: not implemented"; return }

func (m *Model) atContentLimit() bool { _ = "STUB: not implemented"; return false }

func (m *Model) visualLinesForInsert(lines [][]rune) int { _ = "STUB: not implemented"; return 0 }

func (m *Model) mergeLineBelow(row int) { _ = "STUB: not implemented"; return }

func (m *Model) mergeLineAbove(row int) { _ = "STUB: not implemented"; return }

func (m *Model) splitLine(row, col int) { _ = "STUB: not implemented"; return }

func Paste() tea.Msg { _ = "STUB: not implemented"; return *new(tea.Msg) }

func wrap(runes []rune, width int) [][]rune { _ = "STUB: not implemented"; return nil }

//nolint:nestif

func repeatSpaces(n int) []rune { _ = "STUB: not implemented"; return nil }

func numDigits(n int) int { _ = "STUB: not implemented"; return 0 }

func clamp(v, low, high int) int { _ = "STUB: not implemented"; return 0 }

func abs(n int) int { _ = "STUB: not implemented"; return 0 }
