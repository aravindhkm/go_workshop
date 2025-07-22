package concurrency

import (
	"fmt"
	"runtime"
	"time"
)

func heavyTask(id int) {
	start := time.Now()
	count := 0
	for i := 0; i < 1e8; i++ {
		count += i
	}
	fmt.Printf("Goroutine %d finished in %v\n", id, time.Since(start))
}

func OsThreadOne() {
	fmt.Println("Available CPUs:", runtime.NumCPU())
	// fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

	// runtime.GOMAXPROCS(1) // Try changing this to 4

	start := time.Now()
	for i := 1; i <= 4; i++ {
		go heavyTask(i)
	}

	// Wait to let goroutines finish
	time.Sleep(5 * time.Second)
	fmt.Println("Total time:", time.Since(start))
}

// ✅ 1. Goroutines (not OS threads)
// Go uses goroutines to handle concurrency.

// A goroutine is a lightweight, independently executing function managed by the Go runtime.
// Thousands or even millions of goroutines can run concurrently.

// ✅ 2. Go Runtime Scheduler
// The Go runtime includes a scheduler that maps many goroutines onto a smaller number of OS threads.
// It's like a green thread model but more efficient.

// ✅ 3. Multiple OS Threads
// Go does use multiple OS threads when needed.
// By default, it can run goroutines in parallel on multiple CPU cores, not just concurrently.

// ✅ 4. GOMAXPROCS
// The runtime.GOMAXPROCS(n) function sets the maximum number of OS threads that can execute simultaneously.
// By default, it's set to the number of available CPUs.
