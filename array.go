package gokit

import "fmt"

type Array[T comparable] []T

// func (x *Array[T]) At(index int) (v T, ok bool) {
// 	return
// }

func (x *Array[T]) Size() int {
	return len(*x)
}

func (x *Array[T]) Concat(v Array[T]) Array[T] {
	return append(*x, v...)
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
		ok := fn(item)

		if ok {
			v.Append(item)
		}
	}

	return
}

func (x *Array[T]) Find(fn func(v T) bool) (v T, ok bool) {
	for _, item := range *x {
		ok = fn(item)

		if ok {
			v, ok = item, true
			return
		}
	}

	return
}

// func (x *Array[T]) FindLast(fn func(v T) bool) (v T, ok bool) {
// 	for _, item := range *x {
// 		ok = fn(item)

// 		if ok {
// 			v, ok = item, true
// 			return
// 		}
// 	}

// 	return
// }

// func (x *Array[T]) FindIndex(fn func(v T) bool) (v T, ok bool) {
// 	for _, item := range *x {
// 		ok = fn(item)

// 		if ok {
// 			v, ok = item, true
// 			return
// 		}
// 	}

// 	return
// }

// func (x *Array[T]) FindLastIndex(fn func(v T) bool) (v T, ok bool) {
// 	for _, item := range *x {
// 		ok = fn(item)

// 		if ok {
// 			v, ok = item, true
// 			return
// 		}
// 	}

// 	return
// }

// func (x *Array[T]) Reduce(fn func(v T) bool) (v T, ok bool) {
// 	for _, item := range *x {
// 		ok = fn(item)

// 		if ok {
// 			v, ok = item, true
// 			return
// 		}
// 	}

// 	return
// }

// func (x *Array[T]) Slice(fn func(v T) bool) (v T, ok bool) {
// 	for _, item := range *x {
// 		ok = fn(item)

// 		if ok {
// 			v, ok = item, true
// 			return
// 		}
// 	}

// 	return
// }

// func (x *Array[T]) Splice(fn func(v T) bool) (v T, ok bool) {
// 	for _, item := range *x {
// 		ok = fn(item)

// 		if ok {
// 			v, ok = item, true
// 			return
// 		}
// 	}

// 	return
// }

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
	var last int

	if size := x.Size(); size > 0 {
		last, ok = size-1, true
		v, *x = (*x)[last], (*x)[:last]
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

func (x *Array[T]) Join(separator string) (v string) {
	for i, item := range *x {
		if i > 0 {
			v += separator
		}

		v += fmt.Sprint(item)
	}

	return
}
