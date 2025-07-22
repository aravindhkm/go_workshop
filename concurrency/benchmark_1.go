package concurrency

import (
	"fmt"
	"runtime"
	"time"
)

func cpuTask(id int) {
	sum := 0
	for i := 0; i < 1e8; i++ {
		sum += i
		// fmt.Println("sum", id, sum)
	}
}


func runBenchmark(threadCount int) time.Duration {
	runtime.GOMAXPROCS(threadCount)
	start := time.Now()

	done := make(chan bool)
	for i := 0; i < 4; i++ {
		go func(id int) {
			cpuTask(id)
			done <- true
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < 4; i++ {
		<-done
	}

	return time.Since(start)
}

func BenchmarkOne() {
	for _, n := range []int{1, 2, 4, 8} {
		dur := runBenchmark(n)
		fmt.Printf("GOMAXPROCS(%d) => Duration: %v\n", n, dur)
	}
}
