package workerexample

import (
	"fmt"
	"sync"
	"time"
)

func RunUnrestricted(total int) {
	var wg sync.WaitGroup
	fmt.Println("\n▶ Running WITHOUT semaphore (unrestricted goroutines)")

	for i := 1; i <= total; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			fmt.Printf("[Unrestricted] Goroutine %2d STARTED\n", id)

			time.Sleep(1 * time.Second)

			fmt.Printf("[Unrestricted] Goroutine %2d FINISHED\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("✅ Unrestricted run complete")
}

func RunWithSemaphore(total int, maxConcurrent int) {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, maxConcurrent)

	fmt.Println("\n▶ Running WITH semaphore (controlled concurrency)")

	for i := 1; i <= total; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			semaphore <- struct{}{} // Acquire semaphore

			fmt.Printf("[Semaphore   ] Goroutine %2d STARTED | Currently Running: %d\n", id, len(semaphore))

			time.Sleep(1 * time.Second)

			fmt.Printf("[Semaphore   ] Goroutine %2d FINISHED | Currently Running: %d\n", id, len(semaphore))

			<-semaphore // Release semaphore
		}(i)
	}

	wg.Wait()
	fmt.Println("✅ Semaphore-controlled run complete")
}

func CounterSemaphoreOne() {
	totalGoroutines := 10
	// maxConcurrent := 3

	RunUnrestricted(totalGoroutines)
	// time.Sleep(2 * time.Second) // Spacer between runs
	// RunWithSemaphore(totalGoroutines, maxConcurrent)
}
