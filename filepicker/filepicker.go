package filepicker

import (
	"os"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var lastID int64

func nextID() int { _ = "STUB: not implemented"; return 0 }

func New() Model { _ = "STUB: not implemented"; return *new(Model) }

type errorMsg struct {
	err error
}

type readDirMsg struct {
	id      int
	entries []os.DirEntry
}

const (
	marginBottom  = 5
	fileSizeWidth = 7
	paddingLeft   = 2
)

type KeyMap struct {
	GoToTop  key.Binding
	GoToLast key.Binding
	Down     key.Binding
	Up       key.Binding
	PageUp   key.Binding
	PageDown key.Binding
	Back     key.Binding
	Open     key.Binding
	Select   key.Binding
}

func DefaultKeyMap() KeyMap { _ = "STUB: not implemented"; return *new(KeyMap) }

type Styles struct {
	DisabledCursor   lipgloss.Style
	Cursor           lipgloss.Style
	Symlink          lipgloss.Style
	Directory        lipgloss.Style
	File             lipgloss.Style
	DisabledFile     lipgloss.Style
	Permission       lipgloss.Style
	Selected         lipgloss.Style
	DisabledSelected lipgloss.Style
	FileSize         lipgloss.Style
	EmptyDirectory   lipgloss.Style
}

func DefaultStyles() Styles { _ = "STUB: not implemented"; return *new(Styles) }

type Model struct {
	id int

	Path string

	CurrentDirectory string

	AllowedTypes []string

	KeyMap          KeyMap
	files           []os.DirEntry
	ShowPermissions bool
	ShowSize        bool
	ShowHidden      bool
	DirAllowed      bool
	FileAllowed     bool

	FileSelected  string
	selected      int
	selectedStack stack

	minIdx   int
	maxIdx   int
	maxStack stack
	minStack stack

	height     int
	AutoHeight bool

	Cursor string
	Styles Styles
}

type stack struct {
	Push   func(int)
	Pop    func() int
	Length func() int
}

func newStack() stack { _ = "STUB: not implemented"; return *new(stack) }

func (m *Model) pushView(selected, minimum, maximum int) { _ = "STUB: not implemented"; return }

func (m *Model) popView() (int, int, int) { _ = "STUB: not implemented"; return 0, 0, 0 }

func (m Model) readDir(path string, showHidden bool) tea.Cmd {
	_ = "STUB: not implemented"
	return *new(tea.Cmd)
}

func (m *Model) SetHeight(h int) { _ = "STUB: not implemented"; return }

func (m Model) Height() int { _ = "STUB: not implemented"; return 0 }

func (m Model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(Model), *new(tea.Cmd)
}

func (m Model) View() string { _ = "STUB: not implemented"; return "" }

//nolint:gosec

//nolint:nestif

func (m Model) DidSelectFile(msg tea.Msg) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func (m Model) DidSelectDisabledFile(msg tea.Msg) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func (m Model) didSelectFile(msg tea.Msg) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func (m Model) canSelect(file string) bool { _ = "STUB: not implemented"; return false }

func (m Model) HighlightedPath() string { _ = "STUB: not implemented"; return "" }
