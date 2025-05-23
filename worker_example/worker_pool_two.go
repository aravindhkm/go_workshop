package workerexample

import (
	"fmt"
	"sync"
	"time"
)

func workAssign(workerName string, jobId <-chan int, wg *sync.WaitGroup) {
	for getJobId := range jobId {
		fmt.Printf("Worker: %s Job ID: %v Started \n", workerName, getJobId)
		time.Sleep(1 * time.Second)
		fmt.Printf("Worker: %s Job ID: %v Completed \n", workerName, getJobId)
		wg.Done()
	}
}

func WorkerPoolTwo() {
	numOfWorker := []string{"Alice", "Bob", "Eve", "Dave"}
	numOfJobs := 50

	workCh := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < len(numOfWorker); i++ {
		go workAssign(numOfWorker[i], workCh, &wg)
	}

	for jobId := 0; jobId < numOfJobs; jobId++ {
		wg.Add(1)
		workCh <- jobId
	}
	close(workCh)

	wg.Wait() // wait until all jobs are done

	fmt.Println("All work done.")
}
