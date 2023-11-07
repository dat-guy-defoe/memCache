package main

import (
	"github.com/go-co-op/gocron"
	"github.com/stretchr/testify/assert"
	memCache "memCache/internal/cache"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	exitCode := m.Run()

	TestCache(&testing.T{})

	os.Exit(exitCode)
}

func TestCache(t *testing.T) {
	cache := memCache.NewCache()

	// Test setting a value
	cache.Set("key1", "value1", 10)
	assert.Equal(t, "value1", cache.Get("key1"))

	// Test getting a non-existent value
	assert.Nil(t, cache.Get("key2"))

	// Test setting a new value with a longer TTL
	cache.Set("key2", "value2", 20)
	assert.Equal(t, "value2", cache.Get("key2"))

	// Test setting a new value with a shorter TTL
	cache.Set("key3", "value3", 5)
	time.Sleep(10 * time.Second)
	assert.Nil(t, cache.Get("key3"))

	// Test clearing expired values
	cache.Set("key4", "value4", 10)
	time.Sleep(10 * time.Second)
	cache.ClearExpires()
	assert.Nil(t, cache.Get("key4"))

	// Test scheduling clearExpires
	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(5).Seconds().WaitForSchedule().Do(cache.ClearExpires)
	assert.NoError(t, err)
	s.StartAsync()
	time.Sleep(4 * time.Second)
	cache.Set("key", "some value", 5)
	assert.Equal(t, "some value", cache.Get("key"))
	time.Sleep(4 * time.Second)
	cache.Set("key", "some other value", 5)
	assert.Equal(t, "some other value", cache.Get("key"))
	cache.Set("key2", "some other value", 20)
	cache.Set("key3", "some other value", 25)
	cache.Set("key4", "some other value", 3)
	time.Sleep(4 * time.Second)
	assert.NotNil(t, cache.Get("key2"))
	assert.NotNil(t, cache.Get("key3"))
	assert.Nil(t, cache.Get("key4"))
}
