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
func MemLeakOne() {
	fmt.Println("MemLeakOne")
	// leakWithSlice()
}
