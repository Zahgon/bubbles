package key

import "fmt"

type Binding struct {
	keys     []string
	help     Help
	disabled bool
}

type BindingOpt func(*Binding)

func NewBinding(opts ...BindingOpt) Binding { _ = "STUB: not implemented"; return *new(Binding) }

func WithKeys(keys ...string) BindingOpt { _ = "STUB: not implemented"; return *new(BindingOpt) }

func WithHelp(key, desc string) BindingOpt { _ = "STUB: not implemented"; return *new(BindingOpt) }

func WithDisabled() BindingOpt { _ = "STUB: not implemented"; return *new(BindingOpt) }

func (b *Binding) SetKeys(keys ...string) { _ = "STUB: not implemented"; return }

func (b Binding) Keys() []string { _ = "STUB: not implemented"; return nil }

func (b *Binding) SetHelp(key, desc string) { _ = "STUB: not implemented"; return }

func (b Binding) Help() Help { _ = "STUB: not implemented"; return *new(Help) }

func (b Binding) Enabled() bool { _ = "STUB: not implemented"; return false }

func (b *Binding) SetEnabled(v bool) { _ = "STUB: not implemented"; return }

func (b *Binding) Unbind() { _ = "STUB: not implemented"; return }

type Help struct {
	Key  string
	Desc string
}

func Matches[Key fmt.Stringer](k Key, b ...Binding) bool { _ = "STUB: not implemented"; return false }
