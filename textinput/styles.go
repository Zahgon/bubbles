package textinput

import (
	"image/color"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func DefaultStyles(isDark bool) Styles { _ = "STUB: not implemented"; return *new(Styles) }

func DefaultLightStyles() Styles { _ = "STUB: not implemented"; return *new(Styles) }

func DefaultDarkStyles() Styles { _ = "STUB: not implemented"; return *new(Styles) }

type Styles struct {
	Focused StyleState
	Blurred StyleState
	Cursor  CursorStyle
}

type StyleState struct {
	Text        lipgloss.Style
	Placeholder lipgloss.Style
	Suggestion  lipgloss.Style
	Prompt      lipgloss.Style
}

type CursorStyle struct {
	Color color.Color

	Shape tea.CursorShape

	Blink bool

	BlinkSpeed time.Duration
}
