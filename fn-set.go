package gokit

import "fmt"

func Join(sep string, entries ...any) (value string) {
	return JoinFn(sep, entries)
}

func JoinFn(sep string, entries []any, callbacks ...func(v string) string) (value string) {
	for i, entry := range entries {
		data := fmt.Sprint(entry)

		for _, callback := range callbacks {
			if callback == nil {
				continue
			}

			data = callback(data)
		}

		if i > 0 {
			value += sep
		}

		value += data
	}

	return
}

func Split(v any, delim string) (value Array[string]) {
	value = Array[string]{}
	data := String(fmt.Sprint(v))

	data.SplitFn(delim, func(entry String) {
		value = append(value, string(entry))
	})

	return
}

func SplitFn(v any, delim string, callback func(data String)) {
	data := String(fmt.Sprint(v))
	data.SplitFn(delim, callback)
}
