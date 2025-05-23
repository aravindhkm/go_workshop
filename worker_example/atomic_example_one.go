package workerexample

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func atomicRun() {
	var ops atomic.Uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				ops.Add(1)
			}

			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("ops:", ops.Load())
}

func normalRun() {
	var ops int64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				ops += 1
			}

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops:", ops)
}

func normalRunWithLock() {
	var ops int64

	var wg sync.WaitGroup
	mu := sync.Mutex{}

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				mu.Lock()
				ops += 1
				mu.Unlock()
			}

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops:", ops)
}

var runFn = 3

func AtomicExampleOne() {
	if runFn == 1 {
		atomicRun()
	} else if runFn == 2 {
		normalRun()
	} else if runFn == 3 {
		normalRunWithLock()
	}
}
