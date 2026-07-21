package textinput

import (
	"charm.land/bubbles/v2/cursor"
	"charm.land/bubbles/v2/internal/runeutil"
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
)

type (
	pasteMsg    string
	pasteErrMsg struct{ error }
)

type EchoMode int

const (
	EchoNormal EchoMode = iota

	EchoPassword

	EchoNone
)

type ValidateFunc func(string) error

type KeyMap struct {
	CharacterForward        key.Binding
	CharacterBackward       key.Binding
	WordForward             key.Binding
	WordBackward            key.Binding
	DeleteWordBackward      key.Binding
	DeleteWordForward       key.Binding
	DeleteAfterCursor       key.Binding
	DeleteBeforeCursor      key.Binding
	DeleteCharacterBackward key.Binding
	DeleteCharacterForward  key.Binding
	LineStart               key.Binding
	LineEnd                 key.Binding
	Paste                   key.Binding
	AcceptSuggestion        key.Binding
	NextSuggestion          key.Binding
	PrevSuggestion          key.Binding
}

func DefaultKeyMap() KeyMap { _ = "STUB: not implemented"; return *new(KeyMap) }

type Model struct {
	Err error

	Prompt        string
	Placeholder   string
	EchoMode      EchoMode
	EchoCharacter rune

	useVirtualCursor bool

	virtualCursor cursor.Model

	CharLimit int

	styles Styles

	width int

	KeyMap KeyMap

	value []rune

	focus bool

	pos int

	offset      int
	offsetRight int

	Validate ValidateFunc

	rsan runeutil.Sanitizer

	ShowSuggestions bool

	suggestions            [][]rune
	matchedSuggestions     [][]rune
	currentSuggestionIndex int
}

func New() Model { _ = "STUB: not implemented"; return *new(Model) }

func (m Model) VirtualCursor() bool { _ = "STUB: not implemented"; return false }

func (m *Model) SetVirtualCursor(v bool) { _ = "STUB: not implemented"; return }

func (m Model) Styles() Styles { _ = "STUB: not implemented"; return *new(Styles) }

func (m *Model) SetStyles(s Styles) { _ = "STUB: not implemented"; return }

func (m Model) Width() int { _ = "STUB: not implemented"; return 0 }

func (m *Model) SetWidth(w int) { _ = "STUB: not implemented"; return }

func (m *Model) SetValue(s string) { _ = "STUB: not implemented"; return }

func (m *Model) setValueInternal(runes []rune, err error) { _ = "STUB: not implemented"; return }

func (m Model) Value() string { _ = "STUB: not implemented"; return "" }

func (m Model) Position() int { _ = "STUB: not implemented"; return 0 }

func (m *Model) SetCursor(pos int) { _ = "STUB: not implemented"; return }

func (m *Model) CursorStart() { _ = "STUB: not implemented"; return }

func (m *Model) CursorEnd() { _ = "STUB: not implemented"; return }

func (m Model) Focused() bool { _ = "STUB: not implemented"; return false }

func (m *Model) Focus() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m *Model) Blur() { _ = "STUB: not implemented"; return }

func (m *Model) Reset() { _ = "STUB: not implemented"; return }

func (m *Model) SetSuggestions(suggestions []string) { _ = "STUB: not implemented"; return }

func (m *Model) san() runeutil.Sanitizer {
	_ = "STUB: not implemented"
	return *new(runeutil.Sanitizer)
}

func (m *Model) insertRunesFromUserInput(v []rune) { _ = "STUB: not implemented"; return }

func (m *Model) handleOverflow() { _ = "STUB: not implemented"; return }

func (m *Model) deleteBeforeCursor() { _ = "STUB: not implemented"; return }

func (m *Model) deleteAfterCursor() { _ = "STUB: not implemented"; return }

func (m *Model) deleteWordBackward() { _ = "STUB: not implemented"; return }

func (m *Model) deleteWordForward() { _ = "STUB: not implemented"; return }

func (m *Model) wordBackward() { _ = "STUB: not implemented"; return }

func (m *Model) wordForward() { _ = "STUB: not implemented"; return }

func (m Model) echoTransform(v string) string { _ = "STUB: not implemented"; return "" }

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(Model), *new(tea.Cmd)
}

func (m Model) View() string { _ = "STUB: not implemented"; return "" }

//nolint:nestif

func (m Model) promptView() string { _ = "STUB: not implemented"; return "" }

func (m Model) placeholderView() string { _ = "STUB: not implemented"; return "" }

func Blink() tea.Msg { _ = "STUB: not implemented"; return *new(tea.Msg) }

func Paste() tea.Msg { _ = "STUB: not implemented"; return *new(tea.Msg) }

func clamp(v, low, high int) int { _ = "STUB: not implemented"; return 0 }

func (m Model) completionView(offset int) string { _ = "STUB: not implemented"; return "" }

func (m *Model) getSuggestions(sugs [][]rune) []string { _ = "STUB: not implemented"; return nil }

func (m *Model) AvailableSuggestions() []string { _ = "STUB: not implemented"; return nil }

func (m *Model) MatchedSuggestions() []string { _ = "STUB: not implemented"; return nil }

func (m *Model) CurrentSuggestionIndex() int { _ = "STUB: not implemented"; return 0 }

func (m *Model) CurrentSuggestion() string { _ = "STUB: not implemented"; return "" }

func (m *Model) canAcceptSuggestion() bool { _ = "STUB: not implemented"; return false }

func (m *Model) updateSuggestions() { _ = "STUB: not implemented"; return }

func (m *Model) nextSuggestion() { _ = "STUB: not implemented"; return }

func (m *Model) previousSuggestion() { _ = "STUB: not implemented"; return }

func (m Model) validate(v []rune) error { _ = "STUB: not implemented"; return nil }

func (m Model) Cursor() *tea.Cursor { _ = "STUB: not implemented"; return nil }

func (m *Model) updateVirtualCursorStyle() { _ = "STUB: not implemented"; return }

func (m Model) activeStyle() *StyleState { _ = "STUB: not implemented"; return nil }
