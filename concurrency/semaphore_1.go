package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// Semaphore struct
type Semaphore struct {
	permits chan struct{}
}

// NewSemaphore creates a new Semaphore with the given number of permits
func NewSemaphore(maxPermits int) *Semaphore {
	return &Semaphore{permits: make(chan struct{}, maxPermits)}
}

// Acquire a permit from the semaphore
func (s *Semaphore) Acquire() {
	s.permits <- struct{}{}
}

// Release a permit back to the semaphore
func (s *Semaphore) Release() {
	<-s.permits
}

func worker(id int, sem *Semaphore, wg *sync.WaitGroup) {
	defer wg.Done()

	// Acquire a permit before starting work
	sem.Acquire()
	fmt.Printf("Worker %d starting work\n", id)

	// Simulate some work
	time.Sleep(2 * time.Second)
	
	fmt.Printf("Worker %d done working\n", id)

	// Release the permit after finishing work
	sem.Release()
}

func SemaphoreExampleOne() {
	// Create a semaphore with a maximum of 4 permits
	sem := NewSemaphore(4)
	var wg sync.WaitGroup

	// Start 10 workers
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go worker(i, sem, &wg)
	}

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All workers completed")
}
