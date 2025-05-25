package syncexample

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func incrementNonAtomic(wg *sync.WaitGroup, counter *int) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		// This is NOT safe in concurrent environment
		*counter = *counter + 1
	}
}

func incrementAtomic(wg *sync.WaitGroup, counter *int64) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		// Safe atomic increment
		atomic.AddInt64(counter, 1)
	}
}

func Atomic_A() {
	var wg sync.WaitGroup

	// Non-atomic counter
	var counterNonAtomic int
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go incrementNonAtomic(&wg, &counterNonAtomic)
	}
	wg.Wait()
	fmt.Println("Counter without atomic (likely incorrect):", counterNonAtomic)

	// Atomic counter
	var counterAtomic int64
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go incrementAtomic(&wg, &counterAtomic)
	}
	wg.Wait()
	fmt.Println("Counter with atomic (correct):", counterAtomic)
}
