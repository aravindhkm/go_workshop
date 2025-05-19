package task

import (
	"fmt"
	"sync"
)

func printOddTwo(oddCh, evenCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 6 {
		if i%2 != 0 {
			fmt.Println()
			fmt.Println("Odd Block", i)
			<-oddCh // block odd until receive

			// fmt.Println(i)

			fmt.Println("Even Send", i)
			evenCh <- true // release the even block

			// fmt.Println("odd end", i)
			fmt.Println()
		}
	}
}

func printEvenTwo(oddCh, evenCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 6 {
		if i%2 == 0 {
			fmt.Println()
			// fmt.Println(i)

			fmt.Println("Odd Send", i)
			oddCh <- true // release the odd block

			fmt.Println("Even Block", i)
			<-evenCh // block even until receive
			// fmt.Println("even end", i)
			fmt.Println()
		}
	}
}

func AddEvenLinerTwo() {
	oddCh, evenCh := make(chan bool), make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)

	go printOddTwo(oddCh, evenCh, &wg)
	go printEvenTwo(oddCh, evenCh, &wg)

	wg.Wait()

	close(oddCh)
	close(evenCh)

	// <-evenCh
}
