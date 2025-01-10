package selectcase

import "fmt"

func selectExecuteFive(sender chan int, receiver chan string) {
	x, y := 0, 1

	for {
		select {
		case <-receiver:
			{
				fmt.Println("close ch")
				return
			}
		case sender <- x:
			{
				x, y = y, x+y
			}
		}

		// fmt.Println("exec")
	}
}

func ExampleFive() {
	sender := make(chan int)
	receiver := make(chan string)

	go selectExecuteFive(sender, receiver)

	// for i:=0; i<15; i++ {
	// 	fmt.Println("printing", <-sender)
	// }

	for value := range sender {
		fmt.Println("printing", value)

		if value > 1000 {
			receiver <- "close"

			close(receiver)
			close(sender)
		}
	}


	// fmt.Println("Example Five Printing")
}
