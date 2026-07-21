package list

import (
	"cmp"
	"io"
	"time"

	tea "charm.land/bubbletea/v2"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/paginator"
	"charm.land/bubbles/v2/spinner"
	"charm.land/bubbles/v2/textinput"
)

func clamp[T cmp.Ordered](v, low, high T) T { _ = "STUB: not implemented"; return *new(T) }

type Item interface {
	FilterValue() string
}

type ItemDelegate interface {
	Render(w io.Writer, m Model, index int, item Item)

	Height() int

	Spacing() int

	Update(msg tea.Msg, m *Model) tea.Cmd
}

type filteredItem struct {
	index   int
	item    Item
	matches []int
}

type filteredItems []filteredItem

func (f filteredItems) items() []Item { _ = "STUB: not implemented"; return nil }

type FilterMatchesMsg []filteredItem

type FilterFunc func(string, []string) []Rank

type Rank struct {
	Index int

	MatchedIndexes []int
}

func DefaultFilter(term string, targets []string) []Rank { _ = "STUB: not implemented"; return nil }

func UnsortedFilter(term string, targets []string) []Rank { _ = "STUB: not implemented"; return nil }

type statusMessageTimeoutMsg struct{}

type FilterState int

const (
	Unfiltered FilterState = iota
	Filtering
	FilterApplied
)

func (f FilterState) String() string { _ = "STUB: not implemented"; return "" }

type Model struct {
	showTitle        bool
	showFilter       bool
	showStatusBar    bool
	showPagination   bool
	showHelp         bool
	filteringEnabled bool

	itemNameSingular string
	itemNamePlural   string

	Title             string
	Styles            Styles
	InfiniteScrolling bool

	KeyMap KeyMap

	Filter FilterFunc

	disableQuitKeybindings bool

	AdditionalShortHelpKeys func() []key.Binding
	AdditionalFullHelpKeys  func() []key.Binding

	spinner     spinner.Model
	showSpinner bool
	width       int
	height      int
	Paginator   paginator.Model
	cursor      int
	Help        help.Model
	FilterInput textinput.Model
	filterState FilterState

	StatusMessageLifetime time.Duration

	statusMessage      string
	statusMessageTimer *time.Timer

	items []Item

	filteredItems filteredItems

	delegate ItemDelegate
}

func New(items []Item, delegate ItemDelegate, width, height int) Model {
	_ = "STUB: not implemented"
	return *new(Model)
}

func (m *Model) SetFilteringEnabled(v bool) { _ = "STUB: not implemented"; return }

func (m Model) FilteringEnabled() bool { _ = "STUB: not implemented"; return false }

func (m *Model) SetShowTitle(v bool) { _ = "STUB: not implemented"; return }

func (m *Model) SetFilterText(filter string) { _ = "STUB: not implemented"; return }

func (m *Model) SetFilterState(state FilterState) { _ = "STUB: not implemented"; return }

func (m Model) ShowTitle() bool { _ = "STUB: not implemented"; return false }

func (m *Model) SetShowFilter(v bool) { _ = "STUB: not implemented"; return }

func (m Model) ShowFilter() bool { _ = "STUB: not implemented"; return false }

func (m *Model) SetShowStatusBar(v bool) { _ = "STUB: not implemented"; return }

func (m Model) ShowStatusBar() bool { _ = "STUB: not implemented"; return false }

func (m *Model) SetStatusBarItemName(singular, plural string) { _ = "STUB: not implemented"; return }

func (m Model) StatusBarItemName() (string, string) { _ = "STUB: not implemented"; return "", "" }

func (m *Model) SetShowPagination(v bool) { _ = "STUB: not implemented"; return }

func (m *Model) ShowPagination() bool { _ = "STUB: not implemented"; return false }

func (m *Model) SetShowHelp(v bool) { _ = "STUB: not implemented"; return }

func (m Model) ShowHelp() bool { _ = "STUB: not implemented"; return false }

func (m Model) Items() []Item { _ = "STUB: not implemented"; return nil }

func (m *Model) SetItems(i []Item) tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m *Model) Select(index int) { _ = "STUB: not implemented"; return }

func (m *Model) ResetSelected() { _ = "STUB: not implemented"; return }

func (m *Model) ResetFilter() { _ = "STUB: not implemented"; return }

func (m *Model) SetItem(index int, item Item) tea.Cmd {
	_ = "STUB: not implemented"
	return *new(tea.Cmd)
}

func (m *Model) InsertItem(index int, item Item) tea.Cmd {
	_ = "STUB: not implemented"
	return *new(tea.Cmd)
}

func (m *Model) RemoveItem(index int) { _ = "STUB: not implemented"; return }

func (m *Model) SetDelegate(d ItemDelegate) { _ = "STUB: not implemented"; return }

func (m Model) VisibleItems() []Item { _ = "STUB: not implemented"; return nil }

func (m Model) SelectedItem() Item { _ = "STUB: not implemented"; return *new(Item) }

func (m Model) MatchesForItem(index int) []int { _ = "STUB: not implemented"; return nil }

func (m Model) Index() int { _ = "STUB: not implemented"; return 0 }

func (m Model) GlobalIndex() int { _ = "STUB: not implemented"; return 0 }

func (m Model) Cursor() int { _ = "STUB: not implemented"; return 0 }

func (m *Model) CursorUp() { _ = "STUB: not implemented"; return }

func (m *Model) CursorDown() { _ = "STUB: not implemented"; return }

func (m *Model) GoToStart() { _ = "STUB: not implemented"; return }

func (m *Model) GoToEnd() { _ = "STUB: not implemented"; return }

func (m *Model) PrevPage() { _ = "STUB: not implemented"; return }

func (m *Model) NextPage() { _ = "STUB: not implemented"; return }

func (m *Model) maxCursorIndex() int { _ = "STUB: not implemented"; return 0 }

func (m Model) FilterState() FilterState { _ = "STUB: not implemented"; return *new(FilterState) }

func (m Model) FilterValue() string { _ = "STUB: not implemented"; return "" }

func (m Model) SettingFilter() bool { _ = "STUB: not implemented"; return false }

func (m Model) IsFiltered() bool { _ = "STUB: not implemented"; return false }

func (m Model) Width() int { _ = "STUB: not implemented"; return 0 }

func (m Model) Height() int { _ = "STUB: not implemented"; return 0 }

func (m *Model) SetSpinner(spinner spinner.Spinner) { _ = "STUB: not implemented"; return }

func (m *Model) ToggleSpinner() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m *Model) StartSpinner() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m *Model) StopSpinner() { _ = "STUB: not implemented"; return }

func (m *Model) DisableQuitKeybindings() { _ = "STUB: not implemented"; return }

func (m *Model) NewStatusMessage(s string) tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m *Model) SetWidth(v int) { _ = "STUB: not implemented"; return }

func (m *Model) SetHeight(v int) { _ = "STUB: not implemented"; return }

func (m *Model) SetSize(width, height int) { _ = "STUB: not implemented"; return }

func (m *Model) resetFiltering() { _ = "STUB: not implemented"; return }

func (m Model) itemsAsFilterItems() filteredItems {
	_ = "STUB: not implemented"
	return *new(filteredItems)
}

func (m *Model) updateKeybindings() { _ = "STUB: not implemented"; return }

//nolint:exhaustive

func (m *Model) updatePagination() { _ = "STUB: not implemented"; return }

func (m *Model) hideStatusMessage() { _ = "STUB: not implemented"; return }

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(Model), *new(tea.Cmd)
}

func (m *Model) handleBrowsing(msg tea.Msg) tea.Cmd {
	_ = "STUB: not implemented"
	return *new(tea.Cmd)
}

func (m *Model) handleFiltering(msg tea.Msg) tea.Cmd {
	_ = "STUB: not implemented"
	return *new(tea.Cmd)
}

func (m Model) ShortHelp() []key.Binding { _ = "STUB: not implemented"; return nil }

func (m Model) FullHelp() [][]key.Binding { _ = "STUB: not implemented"; return nil }

func (m Model) View() string { _ = "STUB: not implemented"; return "" }

func (m Model) titleView() string { _ = "STUB: not implemented"; return "" }

func (m Model) statusView() string { _ = "STUB: not implemented"; return "" }

//nolint:nestif

//nolint:mnd

func (m Model) paginationView() string { _ = "STUB: not implemented"; return "" }

//nolint:mnd

func (m Model) populatedView() string { _ = "STUB: not implemented"; return "" }

func (m Model) helpView() string { _ = "STUB: not implemented"; return "" }

func (m Model) spinnerView() string { _ = "STUB: not implemented"; return "" }

func filterItems(m Model) tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func insertItemIntoSlice(items []Item, item Item, index int) []Item {
	_ = "STUB: not implemented"
	return nil
}

func removeItemFromSlice(i []Item, index int) []Item { _ = "STUB: not implemented"; return nil }

func removeFilterMatchFromSlice(i []filteredItem, index int) []filteredItem {
	_ = "STUB: not implemented"
	return nil
}

func countEnabledBindings(groups [][]key.Binding) (agg int) { _ = "STUB: not implemented"; return 0 }
