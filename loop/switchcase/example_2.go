package switchcase

import "fmt"

func constraints(args interface{}) {
	switch args.(type) {
		case int: {
			fmt.Println("Interger")
		}
		case string: {
			fmt.Println("String")
		}
	default: {
		fmt.Println("Not found")
	}
	}
}

func ExampleTwo() {
	constraints(5.0)
}