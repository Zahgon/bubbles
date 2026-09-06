package tree

import "charm.land/lipgloss/v2"

type StyleFunc func(children Nodes, i int) lipgloss.Style

type Styles struct {
	TreeStyle lipgloss.Style
	HelpStyle lipgloss.Style

	selectedNodeFunc      StyleFunc
	SelectedNodeStyle     lipgloss.Style
	SelectedNodeStyleFunc StyleFunc

	nodeFunc      StyleFunc
	NodeStyle     lipgloss.Style
	NodeStyleFunc StyleFunc

	rootNodeFunc      StyleFunc
	RootNodeStyle     lipgloss.Style
	RootNodeStyleFunc StyleFunc

	parentNodeFunc      StyleFunc
	ParentNodeStyle     lipgloss.Style
	ParentNodeStyleFunc StyleFunc

	CursorStyle lipgloss.Style

	enumeratorFunc      StyleFunc
	EnumeratorStyle     lipgloss.Style
	EnumeratorStyleFunc StyleFunc

	selectedEnumeratorFunc      StyleFunc
	SelectedEnumeratorStyle     lipgloss.Style
	SelectedEnumeratorStyleFunc StyleFunc

	indenterFunc      StyleFunc
	IndenterStyle     lipgloss.Style
	IndenterStyleFunc StyleFunc

	OpenIndicatorStyle lipgloss.Style
}

func DefaultStyles(isDark bool) (s Styles) { _ = "STUB: not implemented"; return *new(Styles) }

func DefaultDarkStyles() Styles { _ = "STUB: not implemented"; return *new(Styles) }

func DefaultLightStyles() Styles { _ = "STUB: not implemented"; return *new(Styles) }
