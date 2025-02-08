package cache

import (
	"sync"
	"time"
)

type Cache struct {
	data  map[string]CacheItem
	mutex sync.RWMutex
}

type CacheItem struct {
	Value      interface{}
	Expiration int64
}

func NewCache() *Cache {
	return &Cache{data: make(map[string]CacheItem)}
}

// Thread-safe Get operation
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	item, found := c.data[key]
	if !found || time.Now().Unix() > item.Expiration {
		return nil, false
	}
	return item.Value, true
}

// Thread-safe Set operation
func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data[key] = CacheItem{
		Value:      value,
		Expiration: time.Now().Add(duration).Unix(),
	}
}
