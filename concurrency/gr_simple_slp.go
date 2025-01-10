package concurrency

import (
	"fmt"
	"time"
)

func sayFour(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println("print", s)
	}
}

func Gr_Simple_Sleep() {
	go sayFour("world")
	go sayFour("hello")

	// Delay to allow both goroutines to execute
	time.Sleep(1 * time.Second)
}
