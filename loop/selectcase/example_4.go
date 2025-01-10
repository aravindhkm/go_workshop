package selectcase

import "fmt"

func selectExec(worker chan int, quit chan string) {
	curr_state := 0
	for {
		select {
		case worker <- curr_state:
			{
				curr_state++
			}
		case <-quit:
			{
				fmt.Println("quit")
				close(worker)
				return
			}
		}
	}
}

func ExampleFour() {
	worker := make(chan int)
	quit := make(chan string)

	go selectExec(worker,quit)

	for value := range worker {
		fmt.Println("value", value)

		if value == 20 {
			quit <- "done"
		}
	}
}
