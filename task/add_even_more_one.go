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

			evenCh <- true
			if data != lastNum {
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

		// if data == lastNum {
		// 	<-evenCh
		// }
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
