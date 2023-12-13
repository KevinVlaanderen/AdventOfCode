package framework

import "sync"

type SafeCache[T any] struct {
	mu    sync.Mutex
	cache map[Hash]T
}

func NewSafeCache[T any]() *SafeCache[T] {
	return &SafeCache[T]{
		cache: make(map[Hash]T),
	}
}

func (c *SafeCache[T]) Set(key Hash, value T) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = value
}

func (c *SafeCache[T]) Get(key Hash) (T, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.cache[key]
	return value, ok
}
