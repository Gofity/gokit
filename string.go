package gokit

import (
	"strings"
)

type String string

func (x *String) Append(v ...String) {
	for _, entry := range v {
		*x += entry
	}
}

func (x *String) Prepend(v ...String) {
	entries := Array[String](v)
	suffix := entries.Join("")
	*x = suffix + *x
}

func (x String) Split(delim string) (value Array[String]) {
	value = Array[String]{}

	x.SplitFn(delim, func(data String) {
		value.Append(data)
	})

	return
}

func (x String) SplitFn(delim string, callback func(data String)) {
	for {
		index := strings.Index(string(x), delim)

		if index < 0 {
			callback(x)
			break
		}

		callback(x[:index])

		index += len(delim)
		x = x[index:]
	}
}
