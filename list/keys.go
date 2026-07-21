package list

import "charm.land/bubbles/v2/key"

type KeyMap struct {
	CursorUp    key.Binding
	CursorDown  key.Binding
	NextPage    key.Binding
	PrevPage    key.Binding
	GoToStart   key.Binding
	GoToEnd     key.Binding
	Filter      key.Binding
	ClearFilter key.Binding

	CancelWhileFiltering key.Binding
	AcceptWhileFiltering key.Binding

	ShowFullHelp  key.Binding
	CloseFullHelp key.Binding

	Quit key.Binding

	ForceQuit key.Binding
}

func DefaultKeyMap() KeyMap { _ = "STUB: not implemented"; return *new(KeyMap) }
