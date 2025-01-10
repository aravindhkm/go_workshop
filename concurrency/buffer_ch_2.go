package concurrency

import (
	"fmt"
)


func BufferChannelTwo() {
	bufferch := make(chan int,1)
	
	for i:=0; i<15; i++ {
		bufferch <- i
		bufferch <- i + 1
		fmt.Println("Get data from: ", <-bufferch)
		// fmt.Println("Get data from: ", <-bufferch)

	}
	
}