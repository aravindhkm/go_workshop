package interview

import (
	"fmt"
	"sync"
)

func LinearOne() {
	ch := make(chan int)
	pauseCh := make(chan struct{})
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			fmt.Println(<-ch)
			pauseCh <- struct{}{}
		}()
	}

	for value := range 5 {
		ch <- value
		<-pauseCh
	}

	wg.Wait()
	fmt.Println("")
}

// question 
// ch := make(chan int)
// for i := 0; i < 5; i++ {
//     go func() {
//         fmt.Println(<-ch)
//     }()
// }
