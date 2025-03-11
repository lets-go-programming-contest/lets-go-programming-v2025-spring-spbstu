package cache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu      sync.RWMutex
	data    map[string]string
	useSync bool
}

func NewCache(useSync bool) *Cache {
	return &Cache{
		data:    make(map[string]string),
		useSync: useSync,
	}
}

func (c *Cache) Get(key string) (string, bool) {
	if c.useSync {
		c.mu.RLock()
		defer c.mu.RUnlock()
	}
	value, found := c.data[key]
	return value, found
}

func (c *Cache) Set(key, value string) {
	if c.useSync {
		c.mu.Lock()
		defer c.mu.Unlock()
	}
	c.data[key] = value
}

func SimulateExternalFetch(key string) string {
	time.Sleep(100 * time.Millisecond)
	return fmt.Sprintf("Data for %s", key)
}