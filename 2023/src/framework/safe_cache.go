package framework

import "sync"

type SafeCache[K comparable, V any] struct {
	mu    sync.Mutex
	cache map[K]V
}

func NewSafeCache[K comparable, V any]() *SafeCache[K, V] {
	return &SafeCache[K, V]{
		cache: make(map[K]V),
	}
}

func (c *SafeCache[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = value
}

func (c *SafeCache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.cache[key]
	return value, ok
}
