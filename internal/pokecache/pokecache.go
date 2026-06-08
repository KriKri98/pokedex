package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data     map[string]cacheEntry
	mux      sync.RWMutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{interval: interval, data: make(map[string]cacheEntry)}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	val, ok := c.data[key]
	c.mux.RUnlock()
	return val.val, ok
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mux.Lock()
		for key := range c.data {
			if c.data[key].createdAt.Before(time.Now().Add(-c.interval)) {
				delete(c.data, key)
			}
		}
		c.mux.Unlock()
	}
}
