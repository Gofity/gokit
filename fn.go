package gokit

import "reflect"

func Ptr(v any) any {
	vType := reflect.TypeOf(v)

	switch vType.Kind() {
	case reflect.Pointer:
		return v

	default:
		return &v
	}
}
