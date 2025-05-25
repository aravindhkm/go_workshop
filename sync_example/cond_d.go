package syncexample

import (
	"fmt"
	"sync"
)

func Cond_D_One() {
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	cond := sync.NewCond(mu)

	wg.Add(2)

	go func() {
		cond.L.Lock()
		defer cond.L.Unlock()

		fmt.Println("Sending Signel Start")

		for value := range 5 {
			fmt.Println("value", value)
			cond.Signal()
		}

		fmt.Println("Sending Signel End")

		defer wg.Done()
	}()

	go func() {
		cond.L.Lock()
		defer cond.L.Unlock()

		fmt.Println("Waiting Start")
		cond.Wait()
		fmt.Println("Waiting End")

		defer wg.Done()
	}()

	wg.Wait()

	fmt.Println("Exit")
}
