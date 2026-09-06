package spinner

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var lastID int64

func nextID() int { _ = "STUB: not implemented"; return 0 }

type Spinner struct {
	Frames []string
	FPS    time.Duration
}

var (
	Line = Spinner{
		Frames: []string{"|", "/", "-", "\\"},
		FPS:    time.Second / 10, //nolint:mnd
	}
	Dot = Spinner{
		Frames: []string{"⣾ ", "⣽ ", "⣻ ", "⢿ ", "⡿ ", "⣟ ", "⣯ ", "⣷ "},
		FPS:    time.Second / 10, //nolint:mnd
	}
	MiniDot = Spinner{
		Frames: []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
		FPS:    time.Second / 12, //nolint:mnd
	}
	Jump = Spinner{
		Frames: []string{"⢄", "⢂", "⢁", "⡁", "⡈", "⡐", "⡠"},
		FPS:    time.Second / 10, //nolint:mnd
	}
	Pulse = Spinner{
		Frames: []string{"█", "▓", "▒", "░"},
		FPS:    time.Second / 8, //nolint:mnd
	}
	Points = Spinner{
		Frames: []string{"∙∙∙", "●∙∙", "∙●∙", "∙∙●"},
		FPS:    time.Second / 7, //nolint:mnd
	}
	Globe = Spinner{
		Frames: []string{"🌍", "🌎", "🌏"},
		FPS:    time.Second / 4, //nolint:mnd
	}
	Moon = Spinner{
		Frames: []string{"🌑", "🌒", "🌓", "🌔", "🌕", "🌖", "🌗", "🌘"},
		FPS:    time.Second / 8, //nolint:mnd
	}
	Monkey = Spinner{
		Frames: []string{"🙈", "🙉", "🙊"},
		FPS:    time.Second / 3, //nolint:mnd
	}
	Meter = Spinner{
		Frames: []string{
			"▱▱▱",
			"▰▱▱",
			"▰▰▱",
			"▰▰▰",
			"▰▰▱",
			"▰▱▱",
			"▱▱▱",
		},
		FPS: time.Second / 7, //nolint:mnd
	}
	Hamburger = Spinner{
		Frames: []string{"☱", "☲", "☴", "☲"},
		FPS:    time.Second / 3, //nolint:mnd
	}
	Ellipsis = Spinner{
		Frames: []string{"", ".", "..", "..."},
		FPS:    time.Second / 3, //nolint:mnd
	}
)

type Model struct {
	Spinner Spinner

	Style lipgloss.Style

	frame int
	id    int
	tag   int
}

func (m Model) ID() int { _ = "STUB: not implemented"; return 0 }

func New(opts ...Option) Model { _ = "STUB: not implemented"; return *new(Model) }

type TickMsg struct {
	Time time.Time
	tag  int
	ID   int
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(Model), *new(tea.Cmd)
}

func (m Model) View() string { _ = "STUB: not implemented"; return "" }

func (m Model) Tick() tea.Msg { _ = "STUB: not implemented"; return *new(tea.Msg) }

func (m Model) tick(id, tag int) tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

type Option func(*Model)

func WithSpinner(spinner Spinner) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithStyle(style lipgloss.Style) Option { _ = "STUB: not implemented"; return *new(Option) }
