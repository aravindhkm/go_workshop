package interview

import (
	"fmt"
	"sync"
)

func genericCh[I chan int, S chan string, B chan bool, R chan any](iCh I, sCh S, bCh B, resultCh R, wg *sync.WaitGroup) {
	defer wg.Done()
	for intData := range iCh {
		resultCh <- intData
	}

	for intData := range sCh {
		resultCh <- intData
	}

	for intData := range bCh {
		resultCh <- intData
	}
}

func printResultValue(resultCh chan any) {
	for value := range resultCh {
		fmt.Println(value)
	}
	close(resultCh)
}

func GenericChOne() {
	intCh := make(chan int)
	strCh := make(chan string)
	boolCh := make(chan bool)
	resultCh := make(chan any)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go genericCh(intCh, strCh, boolCh, resultCh, wg)
	go printResultValue(resultCh)

	for value := range 5 {
		intCh <- value
	}
	close(intCh)

	for range 5 {
		strCh <- "feinfeih"
	}
	close(strCh)
	
	for range 5 {
		boolCh <- true
	}
	close(boolCh)

	wg.Wait()
	fmt.Println("")
}

// Write a generic function that will combine multiple input
// channels into one output channel
