package interview

import (
	"fmt"
	"sync"
)

// Generic function to combine three channels of different types into one channel of type `any`
func CombineChannels[T1, T2, T3 any](ch1 <-chan T1, ch2 <-chan T2, ch3 <-chan T3, out chan any) {
	wg := sync.WaitGroup{}
	wg.Add(3)

	// Worker to copy from one input channel to output channel
	go func() {
		defer wg.Done()
		for v := range ch1 {
			out <- any(v)
		}
	}()
	go func() {
		defer wg.Done()
		for v := range ch2 {
			out <- v
		}
	}()
	go func() {
		defer wg.Done()
		for v := range ch3 {
			out <- v
		}
	}()

	// Close the output channel once all input channels are drained
	go func() {
		wg.Wait()
		close(out)
	}()
}

func GenericChTwo() {
	intCh := make(chan int)
	stringCh := make(chan string)
	floatCh := make(chan float64)
	outCh := make(chan any)

	// Start feeding data
	go func() {
		defer close(intCh)
		intCh <- 10
		intCh <- 20
	}()

	go func() {
		defer close(stringCh)
		stringCh <- "hello"
		stringCh <- "world"
	}()

	go func() {
		defer close(floatCh)
		floatCh <- 3.14
		floatCh <- 2.71
	}()

	// Combine all into one output channel
	CombineChannels(intCh, stringCh, floatCh, outCh)

	for val := range outCh {
		fmt.Printf("Received: %v (%T)\n", val, val)
	}
}

// any is not a magic type like Object in Java or dynamic in Python
// — it's still statically typed and behaves exactly like interface{}
