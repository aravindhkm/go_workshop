package concurrency

import (
	"fmt"
)


// func printNum(ch chan int) {
// 	for i:=0; i<10; i++ {
// 		// fmt.Println("data printing: ", i)	
// 		ch <- i	
// 	}

// 	close(ch)

// }
// func BufferChannelOne() {
// 	result := make(chan int,5)
// 	go printNum(result)

// 	for value := range result {
// 		fmt.Println("data printing: ", value)	
// 	}
// }

func BufferChannelOne() {
	result := make(chan int,2)
	
	result <- 5
	result <- 1
	// result <- 10

	close(result)

	for value := range result {
		fmt.Println("data printing: ", value)	
	}
}