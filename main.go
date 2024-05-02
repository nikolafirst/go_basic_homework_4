package main

import (
	"fmt"
	"sync"
)

// Cache is a simple in-memory key value store
type Cache struct {
	data map[string]interface{}
	mu   sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.data[key]
	return value, ok
}

func main() {
	cache := NewCache()

	cache.Set("key1", "value1")
	cache.Set("key2", 123)

	val1, ok1 := cache.Get("key1")
	fmt.Println("Value1:", val1, "Exists:", ok1)

	val3, ok3 := cache.Get("key3")
	fmt.Println("Value3:", val3, "Exists:", ok3)
}
