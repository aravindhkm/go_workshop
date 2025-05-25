package syncexample

import (
	"bytes"
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
)

const bufferSize = 1024

// Create a pool for []byte buffers
var bufferPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, bufferSize)
	},
}

func withPool(data string) {
	// Get a buffer from the pool
	buf := bufferPool.Get().([]byte)

	// Use the buffer (e.g., write to it)
	copy(buf, data)

	// Simulate some operation using a bytes.Buffer
	var b bytes.Buffer
	b.Write(buf[:len(data)])
	fmt.Println("withPool:", b.String())

	// Reset and put the buffer back in the pool
	bufferPool.Put(buf)
}

func withoutPool(data string) {
	// Always allocate a new buffer
	buf := make([]byte, bufferSize)

	// Use the buffer (e.g., write to it)
	copy(buf, data)

	var b bytes.Buffer
	b.Write(buf[:len(data)])
	fmt.Println("withoutPool:", b.String())
}

func printMemStats(tag string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s - Alloc = %v KB, TotalAlloc = %v KB, Sys = %v KB, NumGC = %v\n",
		tag, m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
}

// Alloc = currently allocated heap memory

// TotalAlloc = cumulative bytes allocated

// Sys = memory obtained from the OS

// NumGC = number of completed garbage collections

func Pool_A() {
	sample := "Hello, sync.Pool!"

	// Force GC and reset stats before starting
	debug.FreeOSMemory()

	fmt.Println("Calling withoutPool multiple times:")
	printMemStats("Before withoutPool")

	for i := 0; i < 10; i++ {
		withoutPool(sample)
	}

	printMemStats("After withoutPool")

	debug.FreeOSMemory()

	fmt.Println("\nCalling withPool multiple times:")
	printMemStats("Before withPool")

	for i := 0; i < 10; i++ {
		withPool(sample)
	}

	printMemStats("After withPool")
}



// The output will be the same, but behind the scenes:

// withoutPool allocates memory every time.

// withPool reuses memory, which reduces pressure on the garbage collector
// — especially noticeable in high-performance systems.

// pool is goroutine-safe
// pool will not necessarily release the data on the first GC wakeup, but it can release it at any time
// there is no possibility to define and set the pool size
// no need to worry about pool overflow
// You don't need to create a pool everywhere you go; it was designed as a buffer for efficiently sharing common objects, not only within a package but also across multiple packages
// you probably have or will have situations where the need/ability to help GC will be obvious

// https://wundergraph.com/blog/golang-sync-pool#the-allure-of-object-pooling
// https://victoriametrics.com/blog/go-sync-pool/
