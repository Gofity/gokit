package gokit

import (
	"fmt"
	"reflect"
	"slices"
)

type Array[T comparable] []T

func (x *Array[T]) Size() int {
	return len(*x)
}

func (x *Array[T]) LastIndex() int {
	return x.Size() - 1
}

func (x *Array[T]) At(index int) (v T, ok bool) {
	if last := x.LastIndex(); last >= 0 && last >= index {
		return (*x)[index], true
	}

	return
}

func (x *Array[T]) Equal(v Array[T]) bool {
	return slices.Equal(*x, v)
}

func (x *Array[T]) Concat(entries ...Array[T]) (v Array[T]) {
	v = slices.Clone(*x)

	for _, entry := range entries {
		v = append(v, entry...)
	}

	return
}

func (x *Array[T]) Append(v ...T) {
	*x = append(*x, v...)
}

func (x *Array[T]) Prepend(v ...T) {
	*x = append(v, *x...)
}

func (x *Array[T]) Filter(fn func(v T) bool) (v Array[T]) {
	v = Array[T]{}

	for _, item := range *x {
		if ok := fn(item); ok {
			v.Append(item)
		}
	}

	return
}

func (x *Array[T]) Find(fn func(v T) bool) (v T, ok bool) {
	for _, item := range *x {
		if ok = fn(item); ok {
			v, ok = item, true
			return
		}
	}

	return
}

func (x *Array[T]) FindLast(fn func(v T) bool) (v T, ok bool) {
	for i := x.LastIndex(); i >= 0; i-- {
		item := (*x)[i]

		if ok = fn(item); ok {
			v, ok = item, true
			return
		}
	}

	return
}

func (x *Array[T]) FindIndex(fn func(v T) bool) int {
	for i, item := range *x {
		if ok := fn(item); ok {
			return i
		}
	}

	return -1
}

func (x *Array[T]) FindLastIndex(fn func(v T) bool) int {
	for i := x.LastIndex(); i >= 0; i-- {
		item := (*x)[i]

		if ok := fn(item); ok {
			return i
		}
	}

	return -1
}

func (x *Array[T]) Reduce(fn func(accumulator T, v T) T, initial ...T) (v T) {
	if len(initial) > 0 {
		v = initial[0]
	} else {
		vtype := reflect.TypeOf(v)

		if vtype.Kind() == reflect.Pointer {
			v = reflect.New(vtype.Elem()).Interface().(T)
		}
	}

	for _, item := range *x {
		v = fn(v, item)
	}

	return
}

func (x *Array[T]) Sub(start int, count ...int) (v Array[T]) {
	v = Array[T]{}

	if lastIndex := x.LastIndex(); lastIndex >= 0 && lastIndex >= start {
		if len(count) > 0 && count[0] > 0 {
			stop := start + min(count[0], lastIndex)
			return (*x)[start:stop]
		}

		return (*x)[start:]
	}

	return
}

func (x *Array[T]) Slice(start int, end ...int) (v Array[T]) {
	v = Array[T]{}

	if lastIndex := x.LastIndex(); lastIndex >= 0 && lastIndex >= start {
		if len(end) > 0 && lastIndex >= end[0] && end[0] > start {
			return (*x)[start:end[0]]
		}

		return (*x)[start:]
	}

	return
}

func (x *Array[T]) Splice(index, deleteCount int, items ...T) (v Array[T]) {
	v, deleteCount = Array[T]{}, max(deleteCount, 0)

	if lastIndex := x.LastIndex(); lastIndex >= 0 && lastIndex >= index {
		stop := min(index+deleteCount, lastIndex)
		left, right := (*x)[:index], (*x)[stop:]
		return v.Concat(left, items, right)
	}

	return
}

func (x *Array[T]) Map(fn func(v T) T) (v Array[T]) {
	v = Array[T]{}

	for _, item := range *x {
		v.Append(fn(item))
	}

	return
}

func (x *Array[T]) Reverse() (v Array[T]) {
	v = Array[T]{}

	for i := x.Size() - 1; i >= 0; i-- {
		v.Append((*x)[i])
	}

	return
}

func (x *Array[T]) IndexOf(v T) int {
	for i, item := range *x {
		if v == item {
			return i
		}
	}

	return -1
}

func (x *Array[T]) LastIndexOf(v T) int {
	for i := x.Size() - 1; i >= 0; i-- {
		item := (*x)[i]

		if v == item {
			return i
		}
	}

	return -1
}

func (x *Array[T]) Pop() (v T, ok bool) {
	if last := x.LastIndex(); last >= 0 {
		v, *x = (*x)[last], (*x)[:last]
		ok = true
	}

	return
}

func (x *Array[T]) Shift() (v T, ok bool) {
	if x.Size() > 0 {
		v, *x = (*x)[0], (*x)[1:]
		ok = true
	}

	return
}

func (x *Array[T]) Join(separator string) (v String) {
	sep := String(separator)

	for i, item := range *x {
		if i > 0 {
			v += sep
		}

		data := fmt.Sprint(item)
		v += String(data)
	}

	return
}
