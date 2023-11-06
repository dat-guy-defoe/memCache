package main

import (
	"log"
	memCache "memCache/internal/cache"
	"time"
)

func main() {
	cache, err := memCache.NewCache()
	if err != nil {
		log.Fatal(err)
	}

	cache.Set("key", "some value", 5)

	time.Sleep(6 * time.Second)

	value := cache.Get("key")

	log.Println(value)
}
