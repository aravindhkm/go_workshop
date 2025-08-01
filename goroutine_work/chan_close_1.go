package goroutinework

import (
	"fmt"
	"sync"
)

func sendChData(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for value := range 10 {
		ch <- value
	}
	// close(ch) // ✅ Close the channel here
}

func ChanCloseOne() {
	ch := make(chan int)

	wg := &sync.WaitGroup{}

	wg.Add(2)
	go sendChData(ch, wg)
	go sendChData(ch, wg)

	for readCh := range ch {
		fmt.Println(readCh)
	}
	// for range 20 {
	// 	fmt.Println(<-ch)
	// }
	// close(ch)
	// This is wrong because:
	// 1. range ch continues until the channel is closed.
	// 2. So it blocks forever waiting for data if the channel is never closed.
	// 3. You placed close(ch) after the loop — it is never reached until the
	// loop ends, and the loop never ends because the channel is never
	// closed → deadlock.
	wg.Wait()
}
