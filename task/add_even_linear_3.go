package task

import (
	"fmt"
	"sync"
)

func odd(oddCh, evenCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 11; i++ {
		if i%2 != 0 {
			oddCh <- i

			if i != 11 {
				fmt.Println("Even :", <-evenCh)
			} else {
				// <-oddCh
				// if we add this it won't work because 11 data
				// already has been sented. untile repective receiver
				// it not there then it will not allow to execute next code
			}
		}
	}

	close(oddCh) // Signal no more odd numbers

}

func even(oddCh, evenCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 11; i++ {
		if i%2 == 0 {
			fmt.Println("Odd  :", <-oddCh)
			evenCh <- i
		}

		if i == 11 {
			fmt.Println("Odd  :", <-oddCh)
		}
	}
	close(evenCh) // Signal no more even numbers

}

func AddEvenLinerThree() {
	oddCh := make(chan int)
	evenCh := make(chan int)

	wg := &sync.WaitGroup{}

	wg.Add(2)
	go odd(oddCh, evenCh, wg)
	go even(oddCh, evenCh, wg)

	wg.Wait()

	fmt.Println("reverse")
}
