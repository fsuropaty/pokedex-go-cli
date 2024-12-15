package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu sync.Mutex
	c  map[string]cacheEntry
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	cache.c[key] = entry

	cache.mu.Unlock()

}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	entry, exists := cache.c[key]
	cache.mu.Unlock()

	if !exists {
		return nil, false
	}

	return entry.val, true
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		cache.mu.Lock()
		now := time.Now()
		for key, entry := range cache.c {
			if (now.Sub(entry.createdAt)) > interval {
				delete(cache.c, key)
			}
		}
		cache.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		c: make(map[string]cacheEntry),
	}

	go c.reapLoop(interval)
	return c
}
