package syncexample

import (
	"fmt"
	"sync"
	"time"
)

func broadcastOne() {
	var counter int32

	var wg sync.WaitGroup
	var mu sync.Mutex

	cond := sync.NewCond(&mu)

	maxWorkersCount := 10

	wg.Add(maxWorkersCount)

	for i := range maxWorkersCount {
		go func(workerID int) {
			defer wg.Done()

			fmt.Printf("Worker %d performing work\n", workerID)
			time.Sleep(1 * time.Second) // Simulate work

			cond.L.Lock()
			defer cond.L.Unlock()

			counter++

			if counter == int32(maxWorkersCount) {
				fmt.Println("All workers have reached the barrier")
				cond.Broadcast()
			} else {
				fmt.Printf("Worker %d is waiting at the barrier\n", workerID)
				cond.Wait()
			}

			fmt.Printf("Worker %d passed the barrier\n", workerID)
		}(i)
	}

	wg.Wait()
}

func broadcastTwo() {
	wg := &sync.WaitGroup{}

	mu := &sync.Mutex{}
	cond := sync.NewCond(mu)

	maxIndex := 12
	wg.Add(maxIndex)

	runningGoRountine := 0

	for index := 1; index <= maxIndex; index++ {
		go func(worker int) {
			defer wg.Done()

			cond.L.Lock()
			defer cond.L.Unlock()

			runningGoRountine++

			fmt.Println("Worker start the work --------->")

			if maxIndex == runningGoRountine {
				fmt.Println("<------------Close All Work------------->")
				cond.Broadcast()
			} else {
				fmt.Println("Worker waiting id:", worker)
				cond.Wait()
			}

			fmt.Println("Exist", worker)
		}(index)
	}

	wg.Wait()
}

func CondThree() {
	// broadcastOne()
	broadcastTwo()
}


// ### 📈 GitHub Stats

// <p align="center">
//   <img src="https://github-readme-stats.vercel.app/api?username=aravindhkm&show_icons=true&theme=radical" height="160" />
//   <img src="https://github-readme-stats.vercel.app/api/top-langs/?username=aravindhkm&layout=compact&theme=radical" height="160" />
// </p>

// ---