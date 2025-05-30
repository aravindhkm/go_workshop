package task

import (
	"fmt"
	"sync"
)

func printOdd(oddCh, evenCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
			<-oddCh        // block odd until receive
			evenCh <- true // release the even block
		}
	}
}

func printEven(oddCh, evenCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			oddCh <- true // release the odd block
			fmt.Println(i)
			<-evenCh // block even until receive
		}
	}
}

func AddEvenLinerOne() {
	oddCh, evenCh := make(chan bool), make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)

	go printOdd(oddCh, evenCh, &wg)
	go printEven(oddCh, evenCh, &wg)

	wg.Wait()
}

// working program

// func printOdd(oddCh, evenCh chan bool, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 1; i <= 9; i += 2 {
// 		fmt.Println(i)
// 		<-oddCh        		// block odd until receive
// 		evenCh <- true 		// release the even block
// 	}
// }

// func printEven(oddCh, evenCh chan bool, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 2; i <= 10; i += 2 {
// 		oddCh <- true 			// release the odd block
// 		fmt.Println(i)
// 		<-evenCh 				// block even until receive
// 	}
// }

// func AddEvenLiner() {
// 	oddCh, evenCh := make(chan bool), make(chan bool)
// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	go printOdd(oddCh, evenCh, &wg)
// 	go printEven(oddCh, evenCh, &wg)

// 	wg.Wait()
// }
