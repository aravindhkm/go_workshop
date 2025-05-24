package workerexample

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
var currentRunning int

func RunUnrestricted(total int) {
	var wg sync.WaitGroup
	fmt.Println("\n▶ Running WITHOUT semaphore (unrestricted goroutines):\n")

	for i := 1; i <= total; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			mu.Lock()
			currentRunning++
			fmt.Printf("[Unrestricted] Goroutine %2d STARTED | Currently Running: %d\n", id, currentRunning)
			mu.Unlock()

			time.Sleep(1 * time.Second)

			mu.Lock()
			currentRunning--
			fmt.Printf("[Unrestricted] Goroutine %2d FINISHED | Currently Running: %d\n", id, currentRunning)
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println("✅ Unrestricted run complete")
}

func RunWithSemaphore(total int, maxConcurrent int) {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, maxConcurrent)

	fmt.Println("\n▶ Running WITH semaphore (controlled concurrency):\n")

	for i := 1; i <= total; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			semaphore <- struct{}{} // Acquire semaphore

			mu.Lock()
			currentRunning++
			fmt.Printf("[Semaphore   ] Goroutine %2d STARTED | Currently Running: %d\n", id, currentRunning)
			mu.Unlock()

			time.Sleep(1 * time.Second)

			mu.Lock()
			currentRunning--
			fmt.Printf("[Semaphore   ] Goroutine %2d FINISHED | Currently Running: %d\n", id, currentRunning)
			mu.Unlock()

			<-semaphore // Release semaphore
		}(i)
	}

	wg.Wait()
	fmt.Println("✅ Semaphore-controlled run complete")
}

func CounterSemaphoreOne() {
	totalGoroutines := 10
	maxConcurrent := 3

	// RunUnrestricted(totalGoroutines)
	// time.Sleep(2 * time.Second) // Spacer between runs
	RunWithSemaphore(totalGoroutines, maxConcurrent)
}
