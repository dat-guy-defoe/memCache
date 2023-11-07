package cache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	c := NewCache()

	// Test setting a value
	c.Set("key1", "value1", 10)
	if val := c.Get("key1"); val != "value1" {
		t.Errorf("Expected 'value1', but got '%v'", val)
	}

	// Test getting a non-existent value
	if val := c.Get("key2"); val != nil {
		t.Errorf("Expected nil, but got '%v'", val)
	}

	// Test setting a new value with a longer TTL
	c.Set("key2", "value2", 20)
	if val := c.Get("key2"); val != "value2" {
		t.Errorf("Expected 'value2', but got '%v'", val)
	}

	// Test setting a new value with a shorter TTL
	c.Set("key3", "value3", 5)
	time.Sleep(10 * time.Second)
	if val := c.Get("key3"); val != nil {
		t.Errorf("Expected nil, but got '%v'", val)
	}

	// Test clearing expired values
	c.Set("key4", "value4", 10)
	time.Sleep(10 * time.Second)
	c.ClearExpires()
	if val := c.Get("key4"); val != nil {
		t.Errorf("Expected nil, but got '%v'", val)
	}
}
