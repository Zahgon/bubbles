package list

import (
	"charm.land/bubbles/v2/textinput"
	"charm.land/lipgloss/v2"
)

const (
	bullet   = "•"
	ellipsis = "…"
)

type Styles struct {
	TitleBar lipgloss.Style
	Title    lipgloss.Style
	Spinner  lipgloss.Style
	Filter   textinput.Styles

	DefaultFilterCharacterMatch lipgloss.Style

	StatusBar             lipgloss.Style
	StatusEmpty           lipgloss.Style
	StatusBarActiveFilter lipgloss.Style
	StatusBarFilterCount  lipgloss.Style

	NoItems lipgloss.Style

	PaginationStyle lipgloss.Style
	HelpStyle       lipgloss.Style

	ActivePaginationDot   lipgloss.Style
	InactivePaginationDot lipgloss.Style
	ArabicPagination      lipgloss.Style
	DividerDot            lipgloss.Style
}

func DefaultStyles(isDark bool) (s Styles) { _ = "STUB: not implemented"; return *new(Styles) }

//nolint:mnd

//nolint:mnd

//nolint:mnd

//nolint:mnd
