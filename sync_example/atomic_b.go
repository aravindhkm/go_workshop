package syncexample

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func withAtomicB(count *atomic.Int32, wg *sync.WaitGroup) {
	defer wg.Done()

	for range 100000 {
		count.Add(1)
	}
}
func withOutAtomicB(count *int, wg *sync.WaitGroup) {
	defer wg.Done()

	for range 100000 {
		*count++
	}
}

func intPtr(data int) *int {
	return &data
}

func Atomic_B() {
	wg := &sync.WaitGroup{}

	wg.Add(2)
	countForNormal := intPtr(5)
	go withOutAtomicB(countForNormal, wg)
	go withOutAtomicB(countForNormal, wg)

	wg.Add(2)
	countForAtomic := &atomic.Int32{}
	countForAtomic.Add(5)
	go withAtomicB(countForAtomic, wg)
	go withAtomicB(countForAtomic, wg)

	wg.Wait()
	
	fmt.Println("countForNormal", *countForNormal)

	fmt.Println("countForAtomic", countForAtomic.Load())
}
