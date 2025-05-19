package keyword

import (
	"fmt"
	"sync"
)

func printSelect(oddCh chan int, evenCh chan int) {
	for {
		select {
		case even := <-evenCh:
			fmt.Println("even", even)
		case odd := <-oddCh:
			fmt.Println("odd", odd)
		}
	}
}

func SelectOne() {
	oddCh, evenCh := make(chan int), make(chan int)

	wg := &sync.WaitGroup{}

	wg.Add(1)

	go printSelect(oddCh, evenCh)

	for i := range 11 {
		if i%2 == 0 {
			evenCh <- i
		} else {
			oddCh <- i
		}
	}

	wg.Done()
	close(oddCh)
	close(evenCh)

	wg.Wait()

	fmt.Println("Select Exit")
}
