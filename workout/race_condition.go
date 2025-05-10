// package workout

// func RaceCondition() {

// }

package workout

import "fmt"

var counter int

func RaceCondition() {
	go increment()
	go decrement()
	fmt.Println("Counter", counter)
}

func increment() {
	for i := 1; i <= 50; i++ {
		counter++
	}
}

func decrement() {
	for i := 1; i <= 25; i++ {
		counter--
	}
}