package concurrency

import (
	"fmt"
	"sync"
)

func sendMsg(sender chan<- string, wg *sync.WaitGroup) {
	for i:=0; i<=5; i++ {
		sender <- fmt.Sprintf("Sender from %d", i)
	}
	wg.Done()

	close(sender) // Close the channel after sending all messages

}

func receiverMsg(receiver <-chan string,wg *sync.WaitGroup) {
	for receiverMsgContext := range receiver {
		fmt.Println("data", receiverMsgContext)
	}
	wg.Done()
}

func ChannelDirection() {
	wg := sync.WaitGroup{}
	sender := make(chan string)
	// receiver := make(<-chan string)
	wg.Add(2)

	defer wg.Wait()

	go sendMsg(sender,&wg)
	go receiverMsg(sender,&wg)



	fmt.Println("Unidirectional channel")
}

