package interview

import (
	"fmt"
	"sync"
)

func genericSample[T int, B string](intVal T, strVal B) {
	fmt.Println(intVal, strVal)
}

func genericCh[I chan int, S chan string, B chan bool, R chan any](iCh I, sCh S, bCh B, resultCh R, wg *sync.WaitGroup) {
	defer wg.Done()
	// for intData := range iCh {
	// 	resultCh <- intData
	// }

	for range 5 {
		resultCh <- <-iCh
	}
}

func GenericChOne() {
	genericSample(5, "hello world")

	intCh := make(chan int)
	strCh := make(chan string)
	boolCh := make(chan bool)
	resultCh := make(chan any)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go genericCh(intCh, strCh, boolCh, resultCh, wg)

	for value := range 5 {
		intCh <- value
	}

	for range 5 {
		fmt.Println("resultCh", <-resultCh)
	}
	// for result := range resultCh {
	// 	fmt.Println("resultCh", result)
	// }

	// close(resultCh)
	wg.Wait()
	fmt.Println("")
}

// Write a generic function that will combine multiple input
// channels into one output channel
