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

func (x String) Trim(cutset string) String {
	return String(strings.Trim(string(x), cutset))
}

func (x String) TrimFn(fn func(char rune) bool) String {
	return String(strings.TrimFunc(string(x), fn))
}

func (x String) TrimLeft(cutset string) String {
	return String(strings.TrimLeft(string(x), cutset))
}

func (x String) TrimLeftFn(fn func(char rune) bool) String {
	return String(strings.TrimLeftFunc(string(x), fn))
}

func (x String) TrimRight(cutset string) String {
	return String(strings.TrimRight(string(x), cutset))
}

func (x String) TrimRightFn(fn func(char rune) bool) String {
	return String(strings.TrimRightFunc(string(x), fn))
}

func (x String) TrimSpace() String {
	return String(strings.TrimSpace(string(x)))
}

func (x String) TrimPrefix(prefix string) String {
	return String(strings.TrimPrefix(string(x), prefix))
}

func (x String) TrimSuffix(suffix string) String {
	return String(strings.TrimSuffix(string(x), suffix))
}

func (x String) TrimAffix(affix string) String {
	return x.TrimPrefix(affix).TrimSuffix(affix)
}
