package kv

import "sync"

type entity[Key, Value any] struct {
	mu *sync.RWMutex
	mp map[any]Value
}

func (e *entity[Key, Value]) has(key Key) bool {
	e.mu.RLock()
	_, exist := e.mp[key]
	e.mu.RUnlock()
	return exist
}

func (e *entity[Key, Value]) create(key Key, value Value) {
	e.mu.Lock()
	e.mp[key] = value
	e.mu.Unlock()
}

func (e *entity[Key, Value]) update(key Key, value Value) {
	e.mu.RLock()
	e.mp[key] = value
	e.mu.RUnlock()
}

func (e *entity[Key, Value]) delete(key Key) {
	e.mu.Lock()
	delete(e.mp, key)
	e.mu.Unlock()
}

func (e *entity[Key, Value]) r(f func(Key, Value) bool) bool {
	e.mu.RLock()
	defer e.mu.RUnlock()
	for k, v := range e.mp {
		if !f(k.(Key), v) {
			return false
		}
	}
	return true
}

func (e *entity[Key, Value]) query(key Key) (value Value, exist bool) {
	e.mu.RLock()
	value, exist = e.mp[key]
	e.mu.RUnlock()
	return value, exist
}

func newEntity[Key, Value any]() *entity[Key, Value] {
	return &entity[Key, Value]{
		mp: make(map[any]Value),
		mu: &sync.RWMutex{},
	}
}
