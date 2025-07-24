package keyword

import (
	"fmt"
	"sync"
)

func printNum(outCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Recovered from panic:", r)
	// 	}
	// }()

	for i := 1; i < 15; i++ {
		outCh <- i * i
	}
}

//  your original code works fine with i < 10 (or range 9), 
// but if you change the loop to i < 15, you'll run into a deadlock 
// because the main goroutine only reads 9 values from the channel

// But again, this won't help with deadlocks —
// it only works with panics (like writing to a closed channel
// or nil channel).

func RecoverCh() {
	outCh := make(chan int)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go printNum(outCh, wg)

	for range 9 {
		fmt.Println(<-outCh)
	}

	wg.Wait()

	fmt.Println("Recover Run")
}