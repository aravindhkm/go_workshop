package selectcase

import (
	"fmt"
	// "time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			// fmt.Println("x & y", x,y)
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		// default:
		// 	fmt.Println("no value received")
		}

	}
}

func ExampleOne() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			// time.Sleep(time.Second * 2)
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
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