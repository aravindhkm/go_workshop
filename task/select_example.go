package task

import "fmt"


func selectExample(out chan int, quit chan bool) {
	x,y := 0,1

	for {
		select {
		case out <- x:
			{
				x, y = y, x+y
			}
		case <-quit:
			{
				fmt.Println("select closed")
				return
			}
		}
	}
}

func FibonacciExample() {
	data := make(chan int)
	quit := make(chan bool)

	go selectExample(data, quit)	
	
	for i:=0; i<10; i++ {
		fmt.Println(<-data)
	}
	quit <- true
} 