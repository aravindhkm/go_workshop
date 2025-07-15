package interview

import (
	"fmt"
	"time"
)

func fibRun() func() int {
	a, b := 0, 1

	return func() int {
		a, b = b, a+b
		return a
	}
}

func ptr(count int) {
	fmt.Printf("Count %d Time  %s \n", count, time.Now().Format("15:04:05"))
}

func FibTimePing() {
	fibFn := fibRun()

	for count := range 7 {
		n := fibFn()
		fmt.Println(n)
		ptr(count)
		time.Sleep(time.Duration(n) * time.Second)
	}

	fmt.Println("FibTimePing")
}

// // FibonacciGenerator returns a function that yields the next Fibonacci number.
// func FibonacciGenerator() func() int {
// 	a, b := 0, 1
// 	return func() int {
// 		a, b = b, a+b
// 		return a
// 	}
// }

// // FibonacciRateLimiter runs the provided action following Fibonacci delay.
// func FibonacciRateLimiter(action func(int), count int) {
// 	nextFib := FibonacciGenerator()

// 	for i := 1; i <= count; i++ {
// 		delay := nextFib()
// 		fmt.Printf("Waiting %d seconds before action #%d...\n", delay, i)
// 		time.Sleep(time.Duration(delay) * time.Second)
// 		action(i)
// 	}
// }

// func main() {
// 	action := func(i int) {
// 		fmt.Printf("Action #%d executed at %s\n", i, time.Now().Format("15:04:05"))
// 	}

// 	taskCount := 7
// 	FibonacciRateLimiter(action, taskCount)
// }
