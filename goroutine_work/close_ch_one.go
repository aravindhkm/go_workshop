package goroutinework

import "fmt"

func CloseChOne() {
	ch := make(chan int)

	go func() {
		defer close(ch) // Important: signal that sending is done
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()

	// for range close is important
	// for val := range ch { // Stops when channel is closed
	// 	fmt.Println(val)
	// }

	// in this scenario we can use close fn outside goroutine
	// is not mandatory we can n number of loop after close the function it will return zero
	// without close we can't run this
	for range 15 {
		fmt.Println(<-ch)
	}
	// close(ch)

	// If you try to close the channel from outside the sender (e.g., in the main function), there’s a risk:
	// You might close it too early, while the goroutine is still trying to send.
	// Sending to a closed channel causes a panic.
}
