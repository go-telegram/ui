package slider

type Option func(s *Slider)

// OnSelect is a callback function that is called when the user selects a slide
func OnSelect(buttonText string, deleteOnSelect bool, f OnSelectFunc) Option {
	return func(s *Slider) {
		s.selectButtonText = buttonText
		s.deleteOnSelect = deleteOnSelect
		s.onSelect = f
	}
}

// OnCancel is a callback function that is called when the user cancels the slide selection
func OnCancel(buttonText string, deleteOnCancel bool, f OnCancelFunc) Option {
	return func(s *Slider) {
		s.cancelButtonText = buttonText
		s.deleteOnCancel = deleteOnCancel
		s.onCancel = f
	}
}

// OnError is a errors handler
func OnError(f OnErrorFunc) Option {
	return func(s *Slider) {
		s.onError = f
	}
}

// WithPrefix is a keyboard option that sets a prefix for the widget
func WithPrefix(s string) Option {
	return func(w *Slider) {
		w.prefix = s
	}
}
