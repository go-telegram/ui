package dialog

type Option func(d *Dialog)

func Inline() Option {
	return func(d *Dialog) {
		d.inline = true
	}
}

// WithPrefix is a keyboard option that sets a prefix for the widget
func WithPrefix(s string) Option {
	return func(w *Dialog) {
		w.prefix = s
	}
}
