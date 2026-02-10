package slug

import "strings"

type Options struct {
	Delimiter string // Separator. Defaults to `-`
}

func (x *Options) sanitize() {
	x.Delimiter = strings.TrimSpace(x.Delimiter)

	if x.Delimiter == "" {
		x.Delimiter = "-"
	}
}
