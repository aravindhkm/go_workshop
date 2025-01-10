package concurrency

import (
	"fmt"
	"time"
)

// communicate go routine without using any timelock and wait group.

func saytwo(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}
}

func Gr_Simple_Sleep_Two() {
	go saytwo("world")
	saytwo("hello")

}
