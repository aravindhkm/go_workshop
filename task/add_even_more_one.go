package task

import (
	"fmt"
	"sync"
)

const lastNum = 13

func oddChecker(oddCh, evenCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := 1; data <= lastNum; data++ {
		if data%2 != 0 {
			fmt.Println(data)

			// we have to stop it at last value. 
			// if we send it then we need to add respect receiver
			// otherwise we will face block issue
			if data != lastNum {
				evenCh <- true
				<-oddCh
			}
		}
	}
}

func evenChecker(oddCh, evenCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := 1; data <= lastNum; data++ {
		if data%2 == 0 {
			<-evenCh
			fmt.Println(data)
			oddCh <- true
		}
	}
}

func AddEvenMoreOne() {
	oddCh := make(chan bool)
	evenCh := make(chan bool)

	wg := &sync.WaitGroup{}

	wg.Add(2)
	go oddChecker(oddCh, evenCh, wg)
	go evenChecker(oddCh, evenCh, wg)

	wg.Wait()
}
