package task

import "fmt"


func printFibOne(outCh chan int, quit chan bool) {
	a, b := 0,1

	for {
		select {
		case outCh <- a:
			a,b = b+a, a
		case <-quit:
			fmt.Println("closeing")
		}
	}
}

func FibonacciOne() {
	outCh := make(chan int)
	quit := make(chan bool)

	go printFibOne(outCh, quit)

	for range 10 {
		fmt.Println(<-outCh)
	}

	quit <- true

	close(outCh)
	close(quit)
}