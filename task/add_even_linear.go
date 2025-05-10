package task

import (
	"fmt"
	"sync"
)

func printOdd(ch chan int, wg *sync.WaitGroup) {
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			// fmt.Println(i)
			ch <- i
		}
	}

	wg.Done()
}

func printEven(ch chan int, wg *sync.WaitGroup) {
	for i := 1; i <= 10; i++ {
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
	go printOdd(oddCh, &wg)
	go printEven(evenCh, &wg)

	for range [5]int{} {
		fmt.Println(<-oddCh)
		fmt.Println(<-evenCh)
	}
	// close(oddCh)
	// close(evenCh)

	wg.Wait()

}
