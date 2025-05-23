package workerexample

import (
	"fmt"
	"time"
)

func workAssignThree(workerName string, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker: %s Job ID: %v Started \n", workerName, j)
		time.Sleep(time.Second)
		fmt.Printf("Worker: %s Job ID: %v Completed \n", workerName, j)
		results <- j
	}
}

func WorkerPoolThree() {
	numJobs := 10
	worker := [4]string{"Alice", "Bob", "Eve", "Dave"}

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// jobs := make(chan int)
	// results := make(chan int)

	for w := 0; w < len(worker); w++ {
		go workAssignThree(worker[w], jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	for a := 1; a <= numJobs; a++ {
		<-results
	}
	close(jobs)
}
