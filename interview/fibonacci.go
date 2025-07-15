package interview

import (
	"fmt"
	"sync"
)

func fibRunSeries(resultCh chan int, quitCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	a, b := 0, 1
	for {
		select {
		case resultCh <- a:
			a, b = b, a+b
		case <-quitCh:
			fmt.Println("QUIT")
			return
		}
	}
}

func Fibonacci() {
	resultCh := make(chan int)
	quitCh := make(chan struct{})

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go fibRunSeries(resultCh, quitCh, wg)

	for result := range resultCh {
		fmt.Println(result)
		if result > 10 {
			quitCh <- struct{}{}
			close(resultCh)
		}
	}

	close(quitCh)

	wg.Wait()
}
