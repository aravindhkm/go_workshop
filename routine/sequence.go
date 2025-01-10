package routine

import (
	"fmt"
)

func odd(limit int, oddCh chan<- int) {
	for i := 0; i <= limit; i++ {
		if i%2 != 0 {
			oddCh <- i
		}
	}
	close(oddCh)

}

func even(limit int, evenCh chan<- int) {
	for i := 0; i <= limit; i++ {
		if i%2 == 0 {
			evenCh <- i
		}
	}
	close(evenCh)
}

func SequenceMain() {
	limit := 9
	oddCh := make(chan int)
	evenCh := make(chan int)

	go odd(limit, oddCh)
	go even(limit, evenCh)

	var evenNum, oddNum int
	for i := 0; i <= limit; i++ {
		if i%2 == 0 {
			evenNum = <-evenCh
			fmt.Println(evenNum)
		} else {
			oddNum = <-oddCh
			fmt.Println(oddNum)
		}
	}
}
