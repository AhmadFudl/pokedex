package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type (
	Cache struct {
		mux     sync.Mutex
		entries map[string]CacheEntry
	}
	CacheEntry struct {
		createdAt time.Time
		val       any
	}
)

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries: make(map[string]CacheEntry, 4096),
	}

	go func() {
		ticker := time.NewTicker(interval)
		t := time.Time{}
		for {
			select {
			case t = <-ticker.C:
				cache.reap_loop(interval)
			}
			fmt.Println("Tick at: ", t)
		}
	}()

	return cache
}

func (c *Cache) Get(key string) (any, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, ok := c.entries[key]
	return entry.val, ok
}

func (c *Cache) Add(key string, v any) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.entries[key] = CacheEntry{
		createdAt: time.Now(),
		val:       v,
	}
}

func (c *Cache) reap_loop(interval time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.entries {
		if interval <= time.Now().Sub(v.createdAt) {
			delete(c.entries, k)
		}
	}
}
