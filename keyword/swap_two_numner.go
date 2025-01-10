package keyword

import "fmt"

func SwapTwoNumberExecute() {
	   // Original values
	   a := 10
	   b := 20
	   
	   fmt.Println("Before swapping:")
	   fmt.Println("a:", a)
	   fmt.Println("b:", b)
	   
	   // Swapping using a single line
	   a, b = b, a
	   
	   fmt.Println("After swapping:")
	   fmt.Println("a:", a)
	   fmt.Println("b:", b)
}