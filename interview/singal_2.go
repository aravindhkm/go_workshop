package interview

import (
	"fmt"
	"sync"
)

func northSignal(northCh chan struct{}, westCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 12 {
		northCh <- struct{}{}
		<-westCh
		fmt.Println("North Passed")

	}
}

func eastSignal(eastCh chan struct{}, northCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 12 {
		<-northCh
		fmt.Println("North Passed")
		eastCh <- struct{}{}
	}
}

func westSignal(westCh chan struct{}, eastCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 12 {
		<-eastCh
		fmt.Println("East Passed")
		westCh <- struct{}{}
	}
}

func SignalTwo() {
	northCh := make(chan struct{})
	eastCh := make(chan struct{})
	westCh := make(chan struct{})

	wg := &sync.WaitGroup{}

	wg.Add(3)
	go northSignal(northCh, westCh, wg)
	go eastSignal(eastCh, northCh, wg)
	go westSignal(westCh, eastCh, wg)

	wg.Wait()
}
