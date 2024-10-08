package paginator

type Option func(p *Paginator)

// PerPage sets the number of items to be displayed per page.
func PerPage(perPage int) Option {
	return func(p *Paginator) {
		p.perPage = perPage
	}
}

// Separator sets the separator to be used when generating content lines
func Separator(separator string) Option {
	return func(p *Paginator) {
		p.separator = separator
	}
}

// WithCloseButton sets the close button to be displayed
func WithCloseButton(text string) Option {
	return func(p *Paginator) {
		p.closeButton = text
	}
}

// OnError sets the error handler
func OnError(f OnErrorHandler) Option {
	return func(p *Paginator) {
		p.onError = f
	}
}

// WithPrefix is a keyboard option that sets a prefix for the widget
func WithPrefix(s string) Option {
	return func(w *Paginator) {
		w.prefix = s
	}
}

// WithoutEmptyButtons is a keyboard option that hides empty buttons
func WithoutEmptyButtons() Option {
	return func(p *Paginator) {
		p.withoutEmptyButtons = true
	}
}
