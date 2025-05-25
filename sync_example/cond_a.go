package syncexample

import (
	"fmt"
	"sync"
	"time"
)

func CondOne() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	cond := sync.NewCond(&mu)

	wg.Add(2)

	go func() {
		fmt.Println("Goroutine 1 is started")
		defer wg.Done()

		cond.L.Lock()
		defer cond.L.Unlock()

		fmt.Println("Goroutine 1 is waiting for condition")
		cond.Wait()
		fmt.Println("Goroutine 1 met the condition")

		fmt.Println("Goroutine 1 is done")
	}()

	go func() {
		fmt.Println("Goroutine 2 is started")
		defer wg.Done()

		time.Sleep(5 * time.Second) // Simulating some work

		cond.L.Lock()
		defer cond.L.Unlock()

		fmt.Println("Goroutine 2 is signaling condition")
		cond.Signal()
		fmt.Println("Goroutine 2 completed signaling")

		fmt.Println("Goroutine 2 is done")
	}()

	wg.Wait()
}

// The sync.Cond type provides a way to create and manage condition variables.  
// It’s commonly used when one or more goroutines need to wait until some condition 
// becomes true. It has three main methods:

// Wait(): This method causes the calling goroutine to wait until another
// goroutine signals the condition variable. When the goroutine calls Wait(),
// it releases the associated lock and suspends execution until another goroutine
// calls Signal() or Broadcast() on the same sync.Cond variable.

// Signal(): This method wakes up one goroutine waiting on the condition
// variable. If multiple goroutines are waiting, only one of them is awakened.
// The choice of which goroutine gets awakened is arbitrary and not guaranteed.

// Broadcast(): This method wakes up all goroutines waiting for the condition
// variable. When Broadcast() is called, all waiting goroutines are awakened
// and can proceed.

// Note that sync.Cond requires an associated sync.Mutex to synchronize access to
// the condition variable.
// var mu sync.Mutex
// cond := sync.NewCond(&mu)

// cond.Wait() blocks the goroutine and releases the lock until Signal() or Broadcast() is called.
// cond.Signal() wakes one waiting goroutine.
// cond.Broadcast() wakes all waiting goroutines.
