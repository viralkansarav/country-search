package tests

import (
	"sync"
	"testing"
	"time"

	"github.com/viralkansarav/country-search/cache"
)

// testing cache
func TestCache(t *testing.T) {
	c := cache.NewCache()
	c.Set("test", "data", 1*time.Second)
	val, found := c.Get("test")
	if !found || val.(string) != "data" {
		t.Errorf("Cache get failed")
	}
	time.Sleep(2 * time.Second)
	_, found = c.Get("test")
	if found {
		t.Errorf("Cache expiration failed")
	}
}

func TestCacheRaceCondition(t *testing.T) {
	c := cache.NewCache()
	var wg sync.WaitGroup
	key := "testKey"
	value := "testValue"

	// Ensure a value is set before starting concurrent reads
	c.Set(key, value, 5*time.Second)

	// Writer Goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for range 1000 {
			c.Set(key, value, 5*time.Second)
		}
	}()

	// Delay to ensure writer starts before readers
	time.Sleep(10 * time.Millisecond)

	// Reader Goroutines
	for range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				if _, found := c.Get(key); !found {
					t.Errorf("Cache read failed, expected value not found")
				}
			}
		}()
	}

	wg.Wait()
}
