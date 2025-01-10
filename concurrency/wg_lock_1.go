package concurrency

import (
	"fmt"
	"sync"
)

type WgLockType struct {
	Data map[string]int
	sync.Mutex
}

func (d *WgLockType) Set(key string, value int) {
	d.Lock()

	defer d.Unlock()
	d.Data[key] = value
}

func WgLockUnlockExampleOne() {
	state := WgLockType{Data: make(map[string]int)}
	wg := sync.WaitGroup{}
	for i := 0; i < 101; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			state.Set("Hello", i)
		}()
	}

	wg.Wait()
	fmt.Println("Map Value", state.Data["Hello"])
}

// fatal error: concurrent map writes