package gokit

import "reflect"

type Ref[T any] struct{}

func (x *Ref[T]) New() (value T) {
	vType := reflect.TypeOf(value)

	switch vType.Kind() {
	case reflect.Pointer:
		value = reflect.New(vType.Elem()).Interface().(T)

	default:
		value = reflect.Zero(vType).Interface().(T)
	}

	return
}

func (x *Ref[T]) Zero() (value T) {
	vType := reflect.TypeOf(value)
	value = reflect.Zero(vType).Interface().(T)
	return
}
