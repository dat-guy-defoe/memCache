package cache

import (
	"log"
	"time"
)

type Cache struct {
	data    map[string]interface{}
	expires map[string]time.Time
}

func NewCache() *Cache {
	mem := &Cache{
		data:    make(map[string]interface{}),
		expires: make(map[string]time.Time),
	}

	return mem
}

func (c *Cache) ClearExpires() {
	log.Println("run clear", "cache before clear", c.data, c.expires)
	for key, value := range c.expires {
		if value.Before(time.Now()) {
			delete(c.data, key)
			delete(c.expires, key)
		}
	}
	log.Println("cache after clear", c.data, c.expires, "cleared")
}

func (c *Cache) Set(key string, value interface{}, ttl int) {
	c.data[key] = value
	c.expires[key] = time.Now().Add(time.Duration(ttl) * time.Second)
}

func (c *Cache) Get(key string) interface{} {
	if c.expires[key].Before(time.Now()) {
		delete(c.data, key)
		delete(c.expires, key)

		return nil
	}

	v, _ := c.data[key]

	return v
}
