package variadic

import (
	"fmt"
	// "reflect"
)

func add(data ...int) {
	// fmt.Println("data", data, len(data), reflect.TypeOf(data))

	data = append(data, 11)
	for i := 0; i < len(data); i++ {
		fmt.Println("data", data[i])
	}
}

func Run(funcName string) {
	add(1, 2, 3)

	param := []int{4, 5, 6, 7}
	add(param...)
	fmt.Println("variadic")
}
