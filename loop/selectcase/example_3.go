package selectcase

import (
	"fmt"
	// "time"
)

func fibonacci_own(fib_ch chan int,quit chan string) {
	x, y := 0, 1

	for {
		select {
			case fib_ch <- x: 
				x, y = y, y+x
			case <-quit: {
				fmt.Println("quit")
				return
			}
		}
	}

}

func ExampleThree() {
	fib_ch := make(chan int)
	quit := make(chan string)

	go fibonacci_own(fib_ch,quit)

	for i:=0; i<=11; i++ {
		fmt.Println(<-fib_ch)
	}

	quit <- "done"


	fmt.Println("ended")
}



// fibonacci
// 0
// 1
// 1
// 2
// 3
// 5
// 8
// 13
// 21
// 34