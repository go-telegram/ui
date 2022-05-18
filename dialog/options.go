package dialog

type Option func(d *Dialog)

func Inline() Option {
	return func(d *Dialog) {
		d.inline = true
	}
}
