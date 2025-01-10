package switchcase

import "fmt"

func findOddorEven(data int) {
	switch {
		case data%2 == 0: {
			fmt.Println("Even --->", data)
		}
		case data%2 != 0: {
			fmt.Println("Odd  --->",data)
		}
	}
}

func ExampleFour() {

	for i:=0;i<15;i++ {
		findOddorEven(i)
	}
	
	fmt.Println("")
}