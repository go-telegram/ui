package paginator

import (
	"strings"
)

func (p *Paginator) buildText() string {
	if len(p.data) <= p.perPage {
		return strings.Join(p.data, p.separator)
	}

	from := (p.currentPage - 1) * p.perPage
	to := from + p.perPage
	if to > len(p.data) {
		to = len(p.data)
	}

	return strings.Join(p.data[from:to], p.separator)
}
