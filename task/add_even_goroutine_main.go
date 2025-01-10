package task

import (
	"fmt"
)

func oddPrinter(start int,end int, out chan int) {
	for i:=start;i<end;i++ {
		if i % 2 != 0 {
			out <- i
		}
	}

	close(out)
}

func evenPrinter(start int,end int, out chan int) {
	for i:=start;i<end;i++ {
		if i % 2 == 0 {
			out <- i
		}	
	}

	close(out)
}

func PrintAddEvenGoRoutineMain() {
	oddOut := make(chan int)
	evenOut := make(chan int)

	go oddPrinter(1,11,oddOut)
	go evenPrinter(1,11,evenOut)

	for i:=0; i<5; i++ {
		fmt.Println(<-oddOut)
		fmt.Println(<-evenOut)
	}

	fmt.Println("")
}
