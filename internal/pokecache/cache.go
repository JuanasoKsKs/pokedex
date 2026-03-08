package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) Cache {
	mycache := Cache{
		entries: make(map[string]cacheEntry),
		mu:		&sync.Mutex{},
	}
	go mycache.reapLoop(interval)
	return mycache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	result, ok := c.entries[key]
	c.mu.Lock()
	defer c.mu.Unlock()
	if ok {
		return result.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for t := range ticker.C {
		c.mu.Lock()
		oldestPosible := t.Add(-1 * interval)
		for k, entry := range c.entries {
			if entry.createdAt.Before(oldestPosible){
				delete(c.entries, k)
			}
		}
		c.mu.Unlock()
	}
}