package reply

type Option func(kb *ReplyKeyboard)

// WithPrefix is a keyboard option that sets a prefix for the widget
func WithPrefix(s string) Option {
	return func(w *ReplyKeyboard) {
		w.prefix = s
	}
}

// IsSelective is a keyboard option that will show the keyboard only to specific users.
func IsSelective() Option {
	return func(w *ReplyKeyboard) {
		w.selective = true
	}
}

// IsOneTimeKeyboard is a keyboard option that if setted will force the client to hide the keyboard after being used
func IsOneTimeKeyboard() Option {
	return func(w *ReplyKeyboard) {
		w.oneTimeKeyboard = true
	}
}

// ResizableKeyboard is a keyboard option that requests clients to resize the keyboard vertically for optimal fit.
func ResizableKeyboard() Option {
	return func(w *ReplyKeyboard) {
		w.resizeKeyboard = true
	}
}

// IsPersistent is a keyboard option that requests clients to always show the keyboard when the regular keyboard is hidden.
func IsPersistent() Option {
	return func(w *ReplyKeyboard) {
		w.persistent = true
	}
}

// IsPersistent is a keyboard placeholder to be shown in the input field when the keyboard is active
func InputFieldPlaceholder(s string) Option {
	return func(w *ReplyKeyboard) {
		if len(s) > 64 {
			w.inputFieldPlaceholder = ""
		}
		w.inputFieldPlaceholder = s
	}
}
