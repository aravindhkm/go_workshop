package interview

import (
	"fmt"
	"time"
)

func GoRoutineOne() {
	m := map[string]int{}
	go func() {
		m["x"] = 1
	}()
	go func() {
		fmt.Println(m["x"])
	}()
	time.Sleep(time.Second)
}

// 📌 What is the expected output?
// Unpredictable.

// You might get:

// 0 if the read happens before the write.

// 1 if the write happens before the read.

// A crash with a runtime panic (like: fatal error: concurrent map read and map write), especially in later Go versions.

// ❌ What is the problem here?
// This code has a data race.

// ❗ Maps in Go are not safe for concurrent read and write access without synchronization.

// When you read from and write to a map concurrently, without synchronization (e.g., mutex, channel), the program exhibits undefined behavior.

// The Go runtime may crash your program intentionally to protect you from subtle bugs.

// ❌ Problem: Data Race
// Go's built-in map type is not thread-safe.

// In your code:

// One goroutine writes to the map.

// Another goroutine reads from the map.

// These two operations are not synchronized, so they may happen concurrently, causing a data race.

// This can lead to undefined behavior, including:

// Crashes with fatal error: concurrent map read and map write

// Inconsistent results

// | Tool           | Best For               | Thread-Safe | Notes                                  |
// | -------------- | ---------------------- | ----------- | -------------------------------------- |
// | `sync.Mutex`   | General shared data    | ✅           | Manual lock/unlock                     |
// | `sync.RWMutex` | Read-heavy data        | ✅           | Allows concurrent reads                |
// | `sync.Map`     | Concurrent map access  | ✅           | No locking needed, good for cache      |
// | Channels       | Message passing        | ✅           | Go idiomatic, eliminates shared memory |
// | `sync/atomic`  | Simple counters, flags | ✅           | Fastest, limited use cases             |
