package concurrency

import "fmt"

func executeGoRoutine(start, total int, ch chan int, done chan bool) {
	for i := start; i <= total; i++ {
		ch <- i
	}
	done <- true
}

func ParallelismExampleTwo() {
	ch := make(chan int)
	done := make(chan bool, 2) // Make the done channel buffered to prevent deadlock

	go executeGoRoutine(1, 5, ch, done)
	go executeGoRoutine(6, 10, ch, done)

	// work parallelism. print not sequence
	for j := 0; j < 10; j++ {
		fmt.Println(<-ch)
	}

	// not work parallelism. print sequence
	// go func() {
	// 	<-done
	// 	<-done
	// 	close(ch)
	// }()

	// for value := range ch {
	// 	fmt.Println(value)
	// }

}
