package keyword

import "fmt"

type Example struct {
	Name string
	No int
}

var exThree Example

func StructExampleOne() {
	exOne := Example{"Test", 1}

	fmt.Println("exOne", exOne)

	exTwo := Example{
		Name: "Test Two",
		No: 2,
	}

	fmt.Println("exTwo", exTwo)

	exThree = Example{"Test Three", 5}

	fmt.Println("exThree", exThree)

	exFour := Example{Name: "fefegf"}
	fmt.Println("exFour", exFour)

	// error - too few values in struct literal of type ExamplecompilerInvalidStructLit
	// exFive := Example{"fefegf"}
	// fmt.Println("exFour", exFive)
}