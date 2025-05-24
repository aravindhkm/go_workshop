package workerexample

import (
	"fmt"
	"sync"
	"time"
)

func jobAssign(workerName string, jobCh <-chan int, resultCh chan<- int, wg *sync.WaitGroup) {
	for jobID := range jobCh {
		fmt.Printf("Worker: %s Job ID: %v Started\n", workerName, jobID)
		time.Sleep(time.Second) // Simulate some work
		fmt.Printf("Worker: %s Job ID:  ---------> %v Completed\n", workerName, jobID)
		resultCh <- jobID
		wg.Done()
	}
}

func WorkerPoolThree() {
	totalJobs := 15
	workerNames := []string{"Alice", "Bob", "Eve", "Dave"}

	jobCh := make(chan int)
	resultCh := make(chan int)
	var wg sync.WaitGroup

	// Start result collector first
	go func() {
		for result := range resultCh {
			fmt.Println("Ended ------------------------------>", result)
		}
	}()

	// Start workers
	for _, worker := range workerNames {
		go jobAssign(worker, jobCh, resultCh, &wg)
	}

	// Send jobs
	for i := 0; i < totalJobs; i++ {
		wg.Add(1)
		jobCh <- i
	}
	close(jobCh)

	// Wait for all jobs to finish
	wg.Wait()
	close(resultCh)
}
