package keyword

import "fmt"

func PtrTwo() {
	name := "bill"

	namePointer := &name

	fmt.Println(&namePointer)
	printPointer(namePointer)
}

func printPointer(namePointer *string) {
	c := namePointer
	fmt.Println(&namePointer)
	fmt.Println(*c)
}
