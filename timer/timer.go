package timer

import (
	"time"

	tea "charm.land/bubbletea/v2"
)

var lastID int64

func nextID() int { _ = "STUB: not implemented"; return 0 }

type Option func(*Model)

func WithInterval(interval time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

type StartStopMsg struct {
	ID      int
	running bool
}

type TickMsg struct {
	ID int

	Timeout bool

	tag int
}

type TimeoutMsg struct {
	ID int
}

type Model struct {
	Timeout time.Duration

	Interval time.Duration

	id      int
	tag     int
	running bool
}

func New(timeout time.Duration, opts ...Option) Model {
	_ = "STUB: not implemented"
	return *new(Model)
}

func (m Model) ID() int { _ = "STUB: not implemented"; return 0 }

func (m Model) Running() bool { _ = "STUB: not implemented"; return false }

func (m Model) Timedout() bool { _ = "STUB: not implemented"; return false }

func (m Model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(Model), *new(tea.Cmd)
}

func (m Model) View() string { _ = "STUB: not implemented"; return "" }

func (m *Model) Start() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m *Model) Stop() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m *Model) Toggle() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m Model) tick() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m Model) timedout() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m Model) startStop(v bool) tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }
