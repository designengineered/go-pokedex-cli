package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

type Cache struct {
	cacheEntries map[string]cacheEntry
	mutex        sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	// create a new cache pointer and initialize the cache entries map
	c := &Cache{
		cacheEntries: make(map[string]cacheEntry),
	}
	// start a goroutine to reap the cache entries
	go c.ReapLoop(interval)
	// return the cache update
	return c
}

// Reap removes old entries from the cache.
func (c *Cache) Reap(interval time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// get the current time
	loopCountdown := time.Now().UTC().Add(-interval)
	// loop through the cache entries
	for key, value := range c.cacheEntries {
		// if the value is older than the loopCountdown, delete it
		if value.createdAt.Before(loopCountdown) {
			delete(c.cacheEntries, key)
		}
	}
	time.Sleep(interval)
}

func (c *Cache) ReapLoop(interval time.Duration) {
	// start a ticker
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.Reap(interval)
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cacheEntries[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

// Get returns the value associated with the given key.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, ok := c.cacheEntries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}
