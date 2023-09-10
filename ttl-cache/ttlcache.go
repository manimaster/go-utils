package ttlcache

import (
	"sync"
	"time"
)

// Item represents a record in the cache map
type Item struct {
	Value      interface{}
	Expiration int64
}

// Cache represents the TTL cache
type Cache struct {
	items map[string]Item
	mu    sync.RWMutex
}

// NewCache creates a new cache instance
func NewCache() *Cache {
	cache := &Cache{
		items: make(map[string]Item),
	}

	// Start a cleanup goroutine
	go cache.cleanup()

	return cache
}

// Set adds an item to the cache with a TTL
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = Item{
		Value:      value,
		Expiration: time.Now().Add(ttl).UnixNano(),
	}
}

// Get retrieves an item from the cache
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[key]
	if !found {
		return nil, false
	}

	if time.Now().UnixNano() > item.Expiration {
		return nil, false
	}

	return item.Value, true
}

// Delete removes an item from the cache
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
}

// cleanup periodically checks the cache for expired items and removes them
func (c *Cache) cleanup() {
	for {
		time.Sleep(5 * time.Minute)

		c.mu.Lock()
		for key, item := range c.items {
			if time.Now().UnixNano() > item.Expiration {
				delete(c.items, key)
			}
		}
		c.mu.Unlock()
	}
}
