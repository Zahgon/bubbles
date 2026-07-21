package stopwatch

import (
	"time"

	tea "charm.land/bubbletea/v2"
)

var lastID int64

func nextID() int { _ = "STUB: not implemented"; return 0 }

type Option func(*Model)

func WithInterval(interval time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

type TickMsg struct {
	ID  int
	tag int
}

type StartStopMsg struct {
	ID      int
	running bool
}

type ResetMsg struct {
	ID int
}

type Model struct {
	d       time.Duration
	id      int
	tag     int
	running bool

	Interval time.Duration
}

func New(opts ...Option) Model { _ = "STUB: not implemented"; return *new(Model) }

func (m Model) ID() int { _ = "STUB: not implemented"; return 0 }

func (m Model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m Model) Start() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m Model) Stop() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m Model) Toggle() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m Model) Reset() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m Model) Running() bool { _ = "STUB: not implemented"; return false }

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(Model), *new(tea.Cmd)
}

func (m Model) Elapsed() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (m Model) View() string { _ = "STUB: not implemented"; return "" }

func tick(id int, tag int, d time.Duration) tea.Cmd {
	_ = "STUB: not implemented"
	return *new(tea.Cmd)
}
