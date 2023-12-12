package framework

import "sync"

type SafeCache[T any] struct {
	mu    sync.Mutex
	cache map[[64]byte]T
}

func NewSafeCache[T any]() *SafeCache[T] {
	return &SafeCache[T]{
		cache: make(map[[64]byte]T),
	}
}

func (c *SafeCache[T]) Set(key [64]byte, value T) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = value
}

func (c *SafeCache[T]) Get(key [64]byte) (T, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.cache[key]
	return value, ok
}
