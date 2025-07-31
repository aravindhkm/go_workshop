package interview

import (
	"fmt"
	"sync"
)

// chTest reads from iCh and sends each value to resultCh
func chTest(iCh chan int, resultCh chan any, wg *sync.WaitGroup) {
	defer wg.Done()
	for intData := range iCh {
		resultCh <- intData
	}
	close(resultCh) // ✅ Close resultCh after reading all input
}

// issue code
func failDeadLock() {
	intCh := make(chan int)    // ❗️Unbuffered channel
	resultCh := make(chan any) // ❗️Unbuffered channel

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go chTest(intCh, resultCh, wg)

	for value := 0; value < 5; value++ {
		intCh <- value // ✅ Sends data to intCh
	}

	// ❌ This loop causes a deadlock
	for result := range resultCh {
		fmt.Println("resultCh", result)
	}

	close(resultCh) // ❌ UNREACHABLE: Will never be hit because the loop above blocks forever

	wg.Wait()
}

// working Code
func workCh() {
	intCh := make(chan int)
	resultCh := make(chan any)

	var wg sync.WaitGroup
	wg.Add(1)
	go chTest(intCh, resultCh, &wg)

	// ✅ Start a separate goroutine to read from resultCh
	go func() {
		for result := range resultCh {
			fmt.Println("resultCh:", result)
		}
	}()

	// Send data to intCh
	for value := 0; value < 5; value++ {
		intCh <- value
	}
	close(intCh) // ✅ Close intCh to signal completion

	wg.Wait() // Wait for chTest to finish
}

func DeadLockOne() {
	// failDeadLock()
	workCh()
}
