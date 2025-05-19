package concurrency

import "fmt"

func CloseChTwo() {
	ch := make(chan int)

	// Receiver goroutine
	go func() {
		for val := range ch {
			fmt.Println("Received:", val)
		}
		fmt.Println("Channel closed. Exiting receiver.")
	}()

	// Sender
	ch <- 10
	ch <- 20
	ch <- 30

	close(ch) // 🔴 Important: Close the channel to signal no more data

	// Give goroutine time to finish (in real apps use sync.WaitGroup)
	fmt.Scanln()
}
