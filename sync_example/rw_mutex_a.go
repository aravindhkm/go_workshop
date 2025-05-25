package syncexample

import (
	"fmt"
	"sync"
	"time"
)

// Performance comparison between mutex lock and read-write lock

var (
	count int
	//mutex
	countGuard sync.Mutex
	//RWmutex
	countRWGuard sync.RWMutex
)

type mutexMap map[string]string

func (mapData *mutexMap) read() {
	for {
		countGuard.Lock()  // if we comment this we will face this error "fatal error: concurrent map read and map write"
		var _ string = (*mapData)["name"]
		count += 1
		countGuard.Unlock()
	}
}

func (mapData *mutexMap) write() {
	for {
		countGuard.Lock()
		(*mapData)["name"] = "johny"
		count += 1
		time.Sleep(time.Millisecond * 3)
		countGuard.Unlock()
	}
}

type rwMutexMap map[string]string

func (mapData *rwMutexMap) read() {
	for {
		countRWGuard.RLock()
		var _ string = (*mapData)["name"]
		count += 1
		countRWGuard.RUnlock()
	}
}

func (mapData *rwMutexMap) write() {
	for {
		countRWGuard.Lock()
		(*mapData)["name"] = "johny"
		count += 1
		time.Sleep(time.Millisecond * 3)
		countRWGuard.Unlock()
	}
}

func normalMutex() {
	var num int = 3
	mutexMap := mutexMap{"nema": ""}

	for i := 0; i < num; i++ {
		go mutexMap.read()
	}

	for i := 0; i < num; i++ {
		go mutexMap.write()
	}
}

func rwMutex() {
	var num int = 3
	rwMutexMap := rwMutexMap{"nema": ""}

	for i := 0; i < num; i++ {
		go rwMutexMap.read()
	}

	for i := 0; i < num; i++ {
		go rwMutexMap.write()
	}
}

func RW_Mutex_A() {
	normalMutex()
	// rwMutex()

	time.Sleep(time.Second * 3)
	fmt.Printf("Final reads and writes:%d\n", count)
}

// ✅ Use sync.Mutex when reads and writes are roughly equal or simple.

// ✅ Use sync.RWMutex when you have lots of reads and few writes, for better performance.

// Feature	              sync.Mutex	               sync.RWMutex
// Lock type	      Exclusive (all ops)	    Read (RLock) and Write (Lock)
// Read concurrency	        ❌ No	          ✅ Yes (multiple readers allowed)
// Write concurrency	    ❌ No	                ❌ No (only one writer)
// Performance	      Slower for many readers	Faster if many reads, few writes

// The key point is to understand what will happen when use a read-write lock. Here is a blog to explain this:

// Attempting to lock the write lock again when the write lock has been locked will block the current goroutine
// Attempting to lock the read lock again when the write lock is locked will also block the current goroutine
// Attempting to lock the write lock while the read lock is locked will also block the current goroutine
// Attempting to lock the read lock after the read lock has been locked will not block the current goroutine
