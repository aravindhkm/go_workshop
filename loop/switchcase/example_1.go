package switchcase

import "fmt"

func swithchExample(args int) {
	switch cond := args; cond % 10 == 0 {
		case true: {
			fmt.Println("success")
			fmt.Println("value", args)
		}
		case true: {
			fmt.Println("success 2")
			fmt.Println("value", args)
		}
		case false: {
			fmt.Println("fail")
			fmt.Println("value", args)	
		}
		default: {
			fmt.Println("Working fine")
		}
	}

}

func ExampleOne() {
	swithchExample(100)
}	