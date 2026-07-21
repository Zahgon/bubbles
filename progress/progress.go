package progress

import (
	"image/color"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/harmonica"
)

type ColorFunc func(total, current float64) color.Color

var lastID int64

func nextID() int { _ = "STUB: not implemented"; return 0 }

const (
	DefaultFullCharHalfBlock = '▌'

	DefaultFullCharFullBlock = '█'

	DefaultEmptyCharBlock = '░'

	fps              = 60
	defaultWidth     = 40
	defaultFrequency = 18.0
	defaultDamping   = 1.0
)

var (
	defaultBlendStart = lipgloss.Color("#5A56E0")
	defaultBlendEnd   = lipgloss.Color("#EE6FF8")
	defaultFullColor  = lipgloss.Color("#7571F9")
	defaultEmptyColor = lipgloss.Color("#606060")
)

type Option func(*Model)

func WithDefaultBlend() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithColors(colors ...color.Color) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithColorFunc(fn ColorFunc) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithFillCharacters(full rune, empty rune) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithoutPercentage() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithWidth(w int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithSpringOptions(frequency, damping float64) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithScaled(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

type FrameMsg struct {
	id  int
	tag int
}

type Model struct {
	id int

	tag int

	width int

	Full      rune
	FullColor color.Color

	Empty      rune
	EmptyColor color.Color

	ShowPercentage  bool
	PercentFormat   string
	PercentageStyle lipgloss.Style

	spring           harmonica.Spring
	springCustomized bool
	percentShown     float64
	targetPercent    float64
	velocity         float64

	blend []color.Color

	scaleBlend bool

	colorFunc ColorFunc
}

func New(opts ...Option) Model { _ = "STUB: not implemented"; return *new(Model) }

func (m Model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(Model), *new(tea.Cmd)
}

func (m *Model) SetSpringOptions(frequency, damping float64) { _ = "STUB: not implemented"; return }

func (m Model) Percent() float64 { _ = "STUB: not implemented"; return 0 }

func (m *Model) SetPercent(p float64) tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m *Model) IncrPercent(v float64) tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m *Model) DecrPercent(v float64) tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m Model) View() string { _ = "STUB: not implemented"; return "" }

func (m Model) ViewAs(percent float64) string { _ = "STUB: not implemented"; return "" }

func (m *Model) SetWidth(w int) { _ = "STUB: not implemented"; return }

func (m Model) Width() int { _ = "STUB: not implemented"; return 0 }

func (m *Model) nextFrame() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m Model) barView(b *strings.Builder, percent float64, textWidth int) {
	_ = "STUB: not implemented"
	return
}

//nolint:nestif

func (m Model) percentageView(percent float64) string { _ = "STUB: not implemented"; return "" }

//nolint:mnd

func (m *Model) IsAnimating() bool { _ = "STUB: not implemented"; return false }
