package paginator

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
)

type Type int

const (
	Arabic Type = iota
	Dots
)

type KeyMap struct {
	PrevPage key.Binding
	NextPage key.Binding
}

func DefaultKeyMap() KeyMap { _ = "STUB: not implemented"; return *new(KeyMap) }

type Model struct {
	Type Type

	Page int

	PerPage int

	TotalPages int

	ActiveDot string

	InactiveDot string

	ArabicFormat string

	KeyMap KeyMap
}

func (m *Model) SetTotalPages(items int) int { _ = "STUB: not implemented"; return 0 }

func (m Model) ItemsOnPage(totalItems int) int { _ = "STUB: not implemented"; return 0 }

func (m *Model) GetSliceBounds(length int) (start int, end int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (m *Model) PrevPage() { _ = "STUB: not implemented"; return }

func (m *Model) NextPage() { _ = "STUB: not implemented"; return }

func (m Model) OnLastPage() bool { _ = "STUB: not implemented"; return false }

func (m Model) OnFirstPage() bool { _ = "STUB: not implemented"; return false }

type Option func(*Model)

func New(opts ...Option) Model { _ = "STUB: not implemented"; return *new(Model) }

func WithTotalPages(totalPages int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithPerPage(perPage int) Option { _ = "STUB: not implemented"; return *new(Option) }

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(Model), *new(tea.Cmd)
}

func (m Model) View() string {
	_ = "STUB: not implemented"
	//nolint:exhaustive
	return ""
}

func (m Model) dotsView() string { _ = "STUB: not implemented"; return "" }

func (m Model) arabicView() string { _ = "STUB: not implemented"; return "" }
