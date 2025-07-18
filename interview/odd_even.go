package interview

import (
	"fmt"
	"sync"
)

func Odd(oddCh, evenCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i < 11; i++ {
		if i%2 != 0 {
			oddCh <- i
			// fmt.Println(<-evenCh)
		}
	}
}

func Even(oddCh, evenCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i < 11; i++ {
		if i%2 == 0 {
			// fmt.Println(<-oddCh)
			evenCh <- i
		}
	}
}

func OddEvenFind() {
	oddCh := make(chan int)
	evenCh := make(chan int)

	wg := &sync.WaitGroup{}

	wg.Add(2)
	go Odd(oddCh, evenCh, wg)
	go Even(oddCh, evenCh, wg)

	for range 5 {
		fmt.Println(<-oddCh)
		fmt.Println(<-evenCh)
	}

	wg.Wait()

	fmt.Println("Add or Even")
}
