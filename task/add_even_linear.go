package task

import (
	"fmt"
	"sync"
)

func printOdd(ch chan int, wg *sync.WaitGroup) {
	for i := range 10 {
		if i%2 != 0 {
			// fmt.Println(i)
			ch <- i
		}
	}

	wg.Done()
}

func printEven(ch chan int, wg *sync.WaitGroup) {
	for i := range 10 {
		if i%2 == 0 {
			// fmt.Println(i)
			ch <- i
		}
	}

	wg.Done()
}

func AddEvenLiner() {
	oddCh, evenCh := make(chan int), make(chan int)

	wg := sync.WaitGroup{}

	wg.Add(2)
	go printEven(evenCh, &wg)
	go printOdd(oddCh, &wg)

	for range [5]int{} {
		fmt.Println(<-evenCh)
		fmt.Println(<-oddCh)
	}
	// close(oddCh)
	// close(evenCh)

	wg.Wait()

}
