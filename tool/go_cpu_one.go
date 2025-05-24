package tool

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func GoCPU_Debug() {
	runtime.GOMAXPROCS(5) // Use 4 OS threads (simulate multi-core)

	var wg sync.WaitGroup
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers done")
}

// GODEBUG=schedtrace=1000 go run main.go

// SCHED 1008ms: gomaxprocs=8 idleprocs=8 threads=18 spinningthreads=0 needspinning=0 idlethreads=13 runqueue=0 [0 0 0 0 0 0 0 0]
// SCHED 1008ms
// This trace was printed 1008 milliseconds after the program started.

// gomaxprocs=8
// The number of logical processors Go is allowed to use.

// Controlled by runtime.GOMAXPROCS(n).

// In your case, 8 means Go can schedule up to 8 goroutines in parallel (in theory).

// idleprocs=8
// Out of 8 logical processors, 8 are idle at the time of this report.

// That means no goroutines are actively being executed at that moment.

// threads=18
// The total number of OS threads (M) created by the Go runtime.

// These include active, idle, system, and sleeping threads.

// spinningthreads=0
// Number of threads actively spinning, looking for work.

// Threads “spin” before they sleep to avoid sleeping/waking overhead if new work is about to arrive.

// needspinning=0
// Indicates whether the scheduler thinks it needs a spinning thread.

// 0 means the system is calm—no immediate need for more threads to start spinning.

// idlethreads=13
// Number of threads that are not doing anything (idle).

// runqueue=0
// The global run queue length.

// Goroutines ready to run, but not assigned to any processor (P).

// If this number is high, it could indicate goroutines waiting for CPU time.

// [0 0 0 0 0 0 0 0]
// These are the per-P (per logical processor) run queues.

// Each number corresponds to the number of goroutines waiting to run on each logical processor.

// [0 0 0 0 0 0 0 0] means all the per-processor queues are empty at that snapshot.

