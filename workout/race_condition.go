package workout

import (
	"fmt"
	"sync"
)

var counter int

func RaceCondition() {

	wg := sync.WaitGroup{}

	wg.Add(2)
	go increment(&wg)
	go decrement(&wg)

	wg.Wait()
	fmt.Println("Counter", counter)
}

func increment(wg *sync.WaitGroup) {
	for i := 1; i <= 50; i++ {
		counter++
	}
	wg.Done()
}

func decrement(wg *sync.WaitGroup) {
	for i := 1; i <= 25; i++ {
		counter--
	}
	wg.Done()
}

// package workout

// import "fmt"

// var counter int

// func RaceCondition() {
// 	go increment()
// 	go decrement()
// 	fmt.Println("Counter", counter)
// }

// func increment() {
// 	for i := 1; i <= 50; i++ {
// 		counter++
// 	}
// }

// func decrement() {
// 	for i := 1; i <= 25; i++ {
// 		counter--
// 	}
// }
