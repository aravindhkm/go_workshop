package workerexample

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

func CounterSemaphorePkg() {
	var (
		maxConcurrent = int64(3) // Only 3 goroutines can run concurrently
		sem           = semaphore.NewWeighted(maxConcurrent)
		wg            sync.WaitGroup
	)

	ctx := context.Background()
	totalTasks := 10

	for i := 1; i <= totalTasks; i++ {
		wg.Add(1)

		go func(taskID int) {
			defer wg.Done()

			// Acquire a semaphore weight of 1
			if err := sem.Acquire(ctx, 1); err != nil {
				log.Printf("Task %d: failed to acquire semaphore: %v", taskID, err)
				return
			}

			fmt.Printf("Task %d: started\n", taskID)

			// Simulate work
			time.Sleep(2 * time.Second)

			fmt.Printf("Task %d: finished\n", taskID)

			// Release the semaphore
			sem.Release(1)
		}(i)
	}

	wg.Wait()
	fmt.Println("All tasks completed.")
}
