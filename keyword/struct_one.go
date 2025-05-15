package keyword

import "fmt"

type Example struct {
	Name string
	No int
}

func StructExampleOne() {
	exOne := Example{"Test", 1}

	fmt.Println("exOne", exOne)
}