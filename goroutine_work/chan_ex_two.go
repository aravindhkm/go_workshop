package goroutinework

import "fmt"

func ChanExTwo() {
	ch := make(chan int)

	go func() {
		// defer close(ch)
		for value := range 5 {
			ch <- value
		}
	}()

	// it will work
	// for range 10 {
	// 	fmt.Println(<-ch)
	// }

	// for range close is important
	for value := range ch {
		fmt.Println(value)
	}
}
