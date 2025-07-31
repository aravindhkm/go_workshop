package interview

import (
	"fmt"
	"sync"
)

func printSignal(currDir, nextDir chan struct{}, dirName string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 6; i++ {
		<-currDir
		fmt.Println(dirName, "→ Green")
		select {
		case nextDir <- struct{}{}:
		default:
			// Non-blocking send to prevent deadlock if receiver is done
		}

	}
}

func SignalOne() {
	northCh := make(chan struct{})
	eastCh := make(chan struct{})
	westCh := make(chan struct{})

	wg := &sync.WaitGroup{}
	wg.Add(3)

	// Start goroutines
	go printSignal(northCh, eastCh, "NORTH", wg)
	go printSignal(eastCh, westCh, "EAST", wg)
	go printSignal(westCh, northCh, "WEST", wg)

	// Send initial signal to start the chain
	northCh <- struct{}{}

	// Wait for all goroutines to complete
	wg.Wait()

	// Close channels after all goroutines are done
	close(northCh)
	close(eastCh)
	close(westCh)

}
