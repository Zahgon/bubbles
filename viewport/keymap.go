package viewport

import "charm.land/bubbles/v2/key"

type KeyMap struct {
	PageDown     key.Binding
	PageUp       key.Binding
	HalfPageUp   key.Binding
	HalfPageDown key.Binding
	Down         key.Binding
	Up           key.Binding
	Left         key.Binding
	Right        key.Binding
}

func DefaultKeyMap() KeyMap { _ = "STUB: not implemented"; return *new(KeyMap) }
