package gokit

import (
	"reflect"
	"sync"
)

type Ref[T any] = Type[T]

type Type[T any] struct {
	xType reflect.Type
	mutex sync.Mutex
	once  sync.Once
}

func (x *Type[T]) New() (value T) {
	x.once.Do(x.init)

	x.mutex.Lock()
	defer x.mutex.Unlock()

	switch x.xType.Kind() {
	case reflect.Pointer:
		data := reflect.New(x.xType.Elem())
		value = data.Interface().(T)

		// default:
		// 	data := reflect.Zero(x.xType)
		// 	value = data.Interface().(T)
	}

	return
}

func (x *Type[T]) Zero() (value T) {
	x.once.Do(x.init)

	x.mutex.Lock()
	defer x.mutex.Unlock()

	data := reflect.Zero(x.xType)
	value = data.Interface().(T)

	return
}

func (x *Type[T]) Make() (value T) {
	x.once.Do(x.init)

	x.mutex.Lock()
	defer x.mutex.Unlock()

	switch x.xType.Kind() {
	case reflect.Pointer:
		data := reflect.New(x.xType.Elem())
		value = data.Interface().(T)
	}

	return
}

func (x *Type[T]) init() {
	var value T
	x.xType = reflect.TypeOf(value)
}
