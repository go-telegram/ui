package progress

type Option func(p *Progress)

// WithRenderTextFunc sets the render text function.
func WithRenderTextFunc(f RenderTextFunc) Option {
	return func(p *Progress) {
		p.renderTextFunc = f
	}
}

// StartValue sets the start value.
func StartValue(value float64) Option {
	return func(p *Progress) {
		p.value = value
	}
}

// OnError sets the error handler.
func OnError(f OnErrorHandler) Option {
	return func(p *Progress) {
		p.onError = f
	}
}

// WithCancel sets the cancel function.
func WithCancel(buttonText string, deleteOnCancel bool, onCancel OnCancelFunc) Option {
	return func(p *Progress) {
		p.deleteOnCancel = deleteOnCancel
		p.cancelText = buttonText
		p.onCancel = onCancel
	}
}

// WithPrefix is a keyboard option that sets a prefix for the widget
func WithPrefix(s string) Option {
	return func(w *Progress) {
		w.prefix = s
	}
}
