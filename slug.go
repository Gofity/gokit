package gokit

import (
	"bytes"
	"unicode"
)

type xSlug struct{}

func (x *xSlug) Create(text string, delim ...string) string {
	separator := "-"

	if len(delim) > 0 && delim[0] != "" {
		separator = delim[0]
	}

	var buff bytes.Buffer
	var separate bool

	for _, char := range text {
		switch true {
		case x.an(char):
			if separate {
				buff.WriteString(separator)
				separate = false
			}

			char = unicode.ToLower(char)
			buff.WriteRune(char)

		case buff.Len() > 0:
			separate = true
		}
	}

	return buff.String()
}

// Is alphanumeric
func (x *xSlug) an(char rune) bool {
	return unicode.IsLetter(char) || unicode.IsNumber(char)
}

var Slug xSlug
