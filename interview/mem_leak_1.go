package interview

import (
	"fmt"
	"time"
)

var leakySlice []int

// Why it's a leak:
// The leakySlice grows forever and is never cleared or trimmed, causing continuous memory usage.
func leakWithSlice() {
	for {
		leakySlice = append(leakySlice, make([]int, 1000)...)
		time.Sleep(100 * time.Millisecond)
		fmt.Println(leakySlice)
	}
}

type Data struct {
    buffer []byte
}

var globalData []*Data

// Why it's a leak:
// The global slice holds pointers to large allocations, even if they're never used again.
func leakWithPointers() {
    for {
        d := &Data{buffer: make([]byte, 1024*1024)} // 1MB
        globalData = append(globalData, d)
        time.Sleep(100 * time.Millisecond)
    }
}

// ❌ Problem:
// Goroutines that wait forever (e.g., on a channel that never gets a message).
// The goroutines never exit because the channel ch is never closed or written to,
//  and the GC can't clean them up.
func leakWithGoroutines() {
    ch := make(chan int)
    for i := 0; i < 10000; i++ {
        go func() {
            <-ch // Will block forever
        }()
    }
}

func MemLeakOne() {
	fmt.Println("MemLeakOne")
	// leakWithSlice()
}
