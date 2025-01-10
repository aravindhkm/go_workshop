package switchcase

import "fmt"

func switchPrintOddEven(value int, printer chan string) {
	switch {
	case value % 2 == 0:
		{
			printer <- fmt.Sprintf("Number is %d Even", value)
		}
	case value % 2 != 0:
		{
			printer <- fmt.Sprintf("Number is %v Odd", value)
		}
	}
}

func ExampleThree() {
	printer := make(chan string)

	go func() {
		for i:=1; i<12; i++ {
			switchPrintOddEven(i,printer)
		}
	
		close(printer)
	}()

	for result := range printer {
		fmt.Println(result)
	}
}
