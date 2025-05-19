package concurrency

import (
	"fmt"
	"sync"
)

func workerClose(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
	}
}

func CloseChOne() {
	const numJobs = 20
	const numWorkers = 3

	jobs := make(chan int)
	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go workerClose(i, jobs, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	// Close the channel to signal that no more jobs will be sent
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("All jobs processed.")
}
