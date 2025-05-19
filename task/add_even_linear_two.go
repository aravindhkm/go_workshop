package task

import (
	"fmt"
	"sync"
)

func printOddTwo(oddCh, evenCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 10 {
		if i%2 != 0 {
			<-oddCh        // block odd until receive
			fmt.Println(i)
			evenCh <- true // release the even block
		}
	}
}

func printEvenTwo(oddCh, evenCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 10 {
		if i%2 == 0 {
			fmt.Println(i)
			oddCh <- true // release the odd block
			<-evenCh // block even until receive
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
}