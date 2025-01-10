package concurrency

import (
	"fmt"
	"sync"
)

func executeGoRoutineThree(start, total int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; i <= total; i++ {
		fmt.Println(i)
	}
}

func ParallelismExampleThree() {
	var wg sync.WaitGroup
	wg.Add(2)

	go executeGoRoutineThree(1, 5,&wg)
	go executeGoRoutineThree(6, 10,&wg)

	wg.Wait()
}
