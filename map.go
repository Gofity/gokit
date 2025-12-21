package gokit

import (
	"maps"
)

type Map[K comparable, V any] map[K]V

func (x *Map[K, V]) Merge(items ...Map[K, V]) (v Map[K, V]) {
	v = maps.Clone(*x)

	for _, item := range items {
		maps.Copy(v, item)
	}

	return
}

func (x *Map[K, V]) Keys() (v Array[K]) {
	v = Array[K]{}

	for key := range *x {
		v = append(v, key)
	}

	return
}

func (x *Map[K, V]) Values() (v []V) {
	v = []V{}

	for _, value := range *x {
		v = append(v, value)
	}

	return
}
