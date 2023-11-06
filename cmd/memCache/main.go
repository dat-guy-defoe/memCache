package main

import (
	"github.com/go-co-op/gocron"
	"log"
	memCache "memCache/internal/cache"
	"time"
)

func main() {
	cache := memCache.NewCache()

	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(5).Seconds().WaitForSchedule().Do(cache.ClearExpires)
	if err != nil {
		log.Fatal(err)
	}

	cache.Set("key", "some value", 5)

	time.Sleep(4 * time.Second)

	value := cache.Get("key")

	log.Println(value)

	cache.Set("key", "some other value", 5)
	cache.Set("key2", "some other value", 20)
	cache.Set("key3", "some other value", 25)
	cache.Set("key4", "some other value", 3)
	// this keys will be ignored and deleted from cache by clearExpires

	log.Println("do some work")

	s.StartBlocking()
}
