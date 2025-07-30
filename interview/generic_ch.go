package interview

import "fmt"

func genericSample[T int, B string](intVal T, strVal B) {
	fmt.Println(intVal, strVal)
}

func genericCh[I chan int, S chan string, B chan bool, R chan any](iCh I, sCh S, bCh B, resultCh R) {
	for intData := range iCh {
		resultCh <- intData
	}
}

func GenericChOne() {
	genericSample(5, "hello world")

	intCh := make(chan int)
	strCh := make(chan string)
	boolCh := make(chan bool)
	resultCh := make(chan any)

	go genericCh(intCh, strCh, boolCh, resultCh)

	for value := range 5 {
		intCh <- value
	}

	// for result := range

	fmt.Println("")
}

// Write a generic function that will combine multiple input
// channels into one output channel
