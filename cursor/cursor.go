package cursor

import (
	"context"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const defaultBlinkSpeed = time.Millisecond * 530

var lastID int64

func nextID() int { _ = "STUB: not implemented"; return 0 }

type initialBlinkMsg struct{}

type BlinkMsg struct {
	id  int
	tag int
}

type blinkCanceled struct{}

type blinkCtx struct {
	ctx    context.Context
	cancel context.CancelFunc
}

type Mode int

const (
	CursorBlink Mode = iota
	CursorStatic
	CursorHide
)

func (c Mode) String() string { _ = "STUB: not implemented"; return "" }

type Model struct {
	Style lipgloss.Style

	TextStyle lipgloss.Style

	BlinkSpeed time.Duration

	IsBlinked bool

	char string

	id int

	focus bool

	blinkCtx *blinkCtx

	blinkTag int

	mode Mode
}

func New() Model { _ = "STUB: not implemented"; return *new(Model) }

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(Model), *new(tea.Cmd)
}

func (m Model) Mode() Mode { _ = "STUB: not implemented"; return *new(Mode) }

func (m *Model) SetMode(mode Mode) tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m *Model) Blink() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func Blink() tea.Msg { _ = "STUB: not implemented"; return *new(tea.Msg) }

func (m *Model) Focus() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m *Model) Blur() { _ = "STUB: not implemented"; return }

func (m *Model) SetChar(char string) { _ = "STUB: not implemented"; return }

func (m Model) View() string { _ = "STUB: not implemented"; return "" }
