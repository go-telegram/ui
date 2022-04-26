package inline

type Option func(kb *Keyboard)

// NoDeleteAfterClick is a keyboard option that prevents the hide keyboard after click.
func NoDeleteAfterClick() Option {
	return func(kb *Keyboard) {
		kb.deleteAfterClick = false
	}
}

// OnError is a keyboard option that sets a callback function to be called when an error occurs.
func OnError(f func(err error)) Option {
	return func(kb *Keyboard) {
		kb.onError = f
	}
}
