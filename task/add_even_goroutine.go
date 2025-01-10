package task

import (
	"fmt"
	"sync"
)


func printOddNumbers(oddChan, evenChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 { // Iterate only odd numbers
		<-oddChan          // Wait for signal from even goroutine
		fmt.Println(i)     // Print the odd number
		evenChan <- i      // Signal even goroutine
	}
}

func printEvenNumbers(oddChan, evenChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 { // Iterate only even numbers
		<-evenChan         // Wait for signal from odd goroutine
		fmt.Println(i)     // Print the even number
		if i < 10 {        // To avoid sending signal when the sequence is complete
			oddChan <- i   // Signal odd goroutine
		}
	}
}

func PrintAddEvenGoRoutine() {
	oddChan := make(chan int)
	evenChan := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go printOddNumbers(oddChan, evenChan, &wg)
	go printEvenNumbers(oddChan, evenChan, &wg)

	// Start the sequence by sending the first value to oddChan
	oddChan <- 1

	// Wait for both goroutines to finish
	wg.Wait()

	// Close the channels after the goroutines have finished
	close(oddChan)
	close(evenChan)
}
