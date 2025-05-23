package workerexample

import (
	"fmt"
	"sync"
	"time"
)

func workAssignThree(workerName string, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range jobs {
		fmt.Printf("Worker: %s Job ID: %v Started \n", workerName, j)
		time.Sleep(time.Second)
		fmt.Printf("Worker: %s Job ID: %v Completed \n", workerName, j)
		results <- j
	}
}

func WorkerPoolThree() {
	numJobs := 50
	worker := [4]string{"Alice", "Bob", "Eve", "Dave"}
	var wg sync.WaitGroup

	jobs := make(chan int)
	results := make(chan int)

	for w := 0; w < len(worker); w++ {
		wg.Add(1)
		go workAssignThree(worker[w], jobs, results, &wg)
	}

	// Send jobs
	go func() {
		for j := 1; j <= numJobs; j++ {
			jobs <- j
		}
		close(jobs) // Important: close after all jobs are sent
	}()

	// Collect results
	go func() {
		wg.Wait()      // Wait for all workers to finish
		close(results) // Close results channel so we can range over it
	}()

	for res := range results {
		_ = res // Process result if needed
	}
}
