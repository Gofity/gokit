package slug

import (
	"bytes"
	"unicode"
)

func Create(text string, opts Options) string {
	opts.sanitize()

	var buff bytes.Buffer
	var separate bool

	for _, char := range text {
		switch true {
		case alphanumeric(char):
			if separate {
				buff.WriteString(opts.Delimiter)
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

func alphanumeric(char rune) bool {
	return unicode.IsLetter(char) || unicode.IsNumber(char)
}
