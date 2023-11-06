package cache

import (
	"github.com/go-co-op/gocron"
	"time"
)

type Cache struct {
	data    map[string]interface{}
	expires map[string]time.Time
}

func NewCache() (*Cache, error) {
	mem := &Cache{
		data:    make(map[string]interface{}),
		expires: make(map[string]time.Time),
	}

	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(10).Minutes().WaitForSchedule().Do(mem.clearExpires)
	if err != nil {
		return nil, err
	}

	return mem, nil
}

func (c *Cache) clearExpires() {
	for key, value := range c.expires {
		if value.Before(time.Now()) {
			delete(c.data, key)
			delete(c.expires, key)
		}
	}
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
