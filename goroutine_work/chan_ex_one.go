package goroutinework

import (
	"fmt"
	"sync"
)

func sendData(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 5 {
		ch <- true
	}
}

func ChanExOne() {
	chOne := make(chan bool)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go sendData(chOne, wg)

	for range 2 {
		fmt.Println("Print", <-chOne)
	}

	wg.Wait()

	fmt.Println("End")
}


// ❌ Not in unbuffered channels — each send must be matched with a receive.
// ✅ Only in buffered channels, and only up to the buffer size.