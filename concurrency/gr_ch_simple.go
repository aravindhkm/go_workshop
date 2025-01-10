package concurrency

import (
	"fmt"
)

// communicate go routine without using any timelock and wait group.
func say(s string, c chan string) {
	for i := 0; i < 5; i++ {

		if s != "hello" {
			c <- s
		} else {
			fmt.Println(s)
		}
	}
	if s != "hello" {
		close(c) // Close the channel after writing is done
	}
}

func GR_CH_Simple() {
	out := make(chan string)
	go say("world", out)
	say("hello", out)

	for value := range out {
		fmt.Println("-------", value)
	}
}
