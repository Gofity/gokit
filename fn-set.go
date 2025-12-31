package gokit

import "fmt"

func Join(sep string, entries ...any) (value String) {
	return JoinFn(sep, entries)
}

func JoinFn(sep string, entries []any, callbacks ...func(v String) String) (value String) {
	for i, entry := range entries {
		data := String(fmt.Sprint(entry))

		for _, callback := range callbacks {
			if callback == nil {
				continue
			}

			data = callback(data)
		}

		if i > 0 {
			value.Append(String(sep))
		}

		value.Append(data)
	}

	return
}

func Split(v any, delim string) (value Array[String]) {
	data := String(fmt.Sprint(v))
	return data.Split(delim)
}

func SplitFn(v any, delim string, callback func(data String)) {
	data := String(fmt.Sprint(v))
	data.SplitFn(delim, callback)
}
