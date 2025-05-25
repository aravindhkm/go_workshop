package syncexample

import (
	"fmt"
	"sync"
)

func tryRegularMap() {
	regularMap := make(map[int]int)
	var wg sync.WaitGroup
	// mu := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()

			// mu.Lock()
			// defer mu.Unlock()

			// If lock is not there,
			// This can cause: fatal error: concurrent map writes
			regularMap[val] = val

		}(i)
	}

	wg.Wait()

	fmt.Println("Regular map:", regularMap)
}

func trySyncMap() {
	var safeMap sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			safeMap.Store(val, val)
		}(i)
	}

	wg.Wait()

	// Read values from sync.Map
	safeMap.Range(func(key, value any) bool {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
		return true
	})
}

func Map_A() {
	// fmt.Println("=== Without sync.Map (will panic) ===")
	// tryRegularMap()

	fmt.Println("\n=== With sync.Map (safe) ===")
	trySyncMap()
}

// https://victoriametrics.com/blog/go-sync-map/index.html
