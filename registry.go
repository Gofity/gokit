package gokit

import "sync"

type Registry[K comparable, T any] struct {
	mutex  sync.Mutex
	mapper map[K]T
}

func (x *Registry[K, T]) Add(key K, value T, override ...bool) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	x.initialize()

	shouldOverride := len(override) > 0 && override[0]

	if _, ok := x.mapper[key]; ok && !shouldOverride {
		return
	}

	x.mapper[key] = value
}

func (x *Registry[K, T]) Get(key K) (value T) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	x.initialize()

	if v, ok := x.mapper[key]; ok {
		value = v
	}

	return
}

func (x *Registry[K, T]) Has(key K) (exists bool) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	x.initialize()

	_, exists = x.mapper[key]

	return
}

func (x *Registry[K, T]) Remove(key K) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	delete(x.mapper, key)
}

// ================================

func (x *Registry[K, T]) initialize() {
	if x.mapper == nil {
		x.mapper = make(map[K]T)
	}
}
