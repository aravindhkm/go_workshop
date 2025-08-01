package goroutinework

import "fmt"

func sendChData(ch chan int) {
	for value := range 80 {
		ch <- value
	}
	close(ch)
}

func ChanCloseOne() {
	ch := make(chan int)

	go sendChData(ch)

	for readCh := range ch {
		fmt.Println(readCh)
	}
}
