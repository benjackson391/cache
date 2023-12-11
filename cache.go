package cache

import (
	"fmt"
	"sync"
	"time"
)

type Value struct {
	value interface{}
	ttl   *time.Time
}

type Cache struct {
	mu   sync.RWMutex
	data map[string]Value
}

func New() *Cache {
	cache := &Cache{
		data: make(map[string]Value),
	}

	go func() {
		fmt.Println("Starting cache cleaner " + time.Now().String())
		for {
			cache.mu.Lock()
			for key, value := range cache.data {
				if value.ttl != nil && time.Since(*value.ttl) > 0 {
					cache.Delete(key)
				}
			}
			cache.mu.Unlock()
			fmt.Println("Cache cleaner finished " + time.Now().String())
			time.Sleep(time.Microsecond * 100)
		}
	}()

	return cache
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	expiration_date := time.Now().Add(ttl)

	c.data[key] = Value{
		value: value,
		ttl:   &expiration_date,
	}
}

func (c *Cache) Get(key string) interface{} {
	return c.data[key].value
}

func (c *Cache) Delete(key string) {
	delete(c.data, key)
}
