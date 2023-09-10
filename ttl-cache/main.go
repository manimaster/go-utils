package main

import (
	"fmt"
	"time"
	"ttlcache"
)

func main() {
	cache := ttlcache.NewCache()

	cache.Set("name", "John Doe", 5*time.Second)

	value, found := cache.Get("name")
	if found {
		fmt.Println("Found:", value)
	} else {
		fmt.Println("Not found")
	}

	time.Sleep(6 * time.Second)

	value, found = cache.Get("name")
	if found {
		fmt.Println("Found:", value)
	} else {
		fmt.Println("Not found")
	}
}
