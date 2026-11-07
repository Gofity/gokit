package gokit

import "reflect"

type CopyHander[T any] func(v T) T

func Copy[T any](data T, handlers ...CopyHander[T]) (v T) {
	xValue := reflect.ValueOf(data)

	switch xValue.Type().Kind() {
	case reflect.Pointer:
		value := reflect.New(xValue.Elem().Type())
		value.Elem().Set(xValue.Elem())

		v = value.Interface().(T)
	default:
		v = data
	}

	// Call merge handlers

	for _, handler := range handlers {
		v = handler(v)
	}

	return
}

func Make[T any]() (value T) {
	xType := reflect.TypeOf(value)

	if xType.Kind() == reflect.Pointer {
		data := reflect.New(xType.Elem())
		value = data.Interface().(T)
	}

	return
}

func Ptr(v any) any {
	vType := reflect.TypeOf(v)

	switch vType.Kind() {
	case reflect.Pointer:
		return v

	default:
		return &v
	}
}

func IsNil(v any) bool {
	xValue := reflect.ValueOf(v)

	switch xValue.Type().Kind() {
	case reflect.Chan,
		reflect.Func,
		reflect.Interface,
		reflect.Map,
		reflect.Pointer,
		reflect.UnsafePointer,
		reflect.Slice:
		return xValue.IsNil()
	}

	return false
}
