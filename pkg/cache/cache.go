package cache

import (
	"sync"
	"time"
)

const (
	NoExpiry = time.Duration(0)
)

// Cache is a simple generic cache that can be used to store any type of data
type Cache[T any] struct {
	mu      sync.Mutex
	data    map[string]item[T]
	options options
}

type item[T any] struct {
	value  T
	expiry int64
}

// New returns a new cache of the given type
func New[T any](opts ...Option) *Cache[T] {
	def := defaultOptions()
	for _, o := range opts {
		o(&def)
	}

	return &Cache[T]{
		mu:      sync.Mutex{},
		data:    make(map[string]item[T]),
		options: def,
	}
}

// Get returns the value for the given key if it exists and has not expired
func (c *Cache[T]) Get(key string) (T, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var v T
	i, ok := c.data[key]
	if !ok {
		return v, false
	}

	if i.expiry != 0 && i.expiry < time.Now().UnixNano() {
		delete(c.data, key)
		return v, false
	}

	return i.value, true
}

// Set sets the value for the given key with the default expiry
func (c *Cache[T]) Set(key string, value T) {
	c.SetWithExpiry(key, value, c.options.expiry)
}

// SetWithExpiry sets the value for the given key with an expiry
func (c *Cache[T]) SetWithExpiry(key string, value T, expiry time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if expiry == NoExpiry {
		c.data[key] = item[T]{
			value:  value,
			expiry: expiry.Nanoseconds(),
		}
		return
	}

	c.data[key] = item[T]{
		value:  value,
		expiry: time.Now().Add(expiry).UnixNano(),
	}
}

// Delete deletes the value for the given key
func (c *Cache[T]) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, key)
}

// Clear clears the cache
func (c *Cache[T]) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data = make(map[string]item[T])
}

// Len returns the number of items in the cache
func (c *Cache[T]) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return len(c.data)
}

// Keys returns the keys in the cache
func (c *Cache[T]) Keys() []string {
	c.mu.Lock()
	defer c.mu.Unlock()

	keys := make([]string, 0, len(c.data))
	for k := range c.data {
		keys = append(keys, k)
	}

	return keys
}

// Flush removes expired items from the cache
func (c *Cache[T]) Flush() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now().UnixNano()

	for k := range c.data {
		if c.data[k].expiry == 0 || now < c.data[k].expiry {
			continue
		}
		delete(c.data, k)
	}
}
