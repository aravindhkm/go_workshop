package generics

import (
	"fmt"
	"reflect"
)

type ITF interface{}

func GenericExampleOne[T ITF](data []T) {
	for index, value := range data {
		fmt.Println("index", index)
		fmt.Println("reflect type", reflect.TypeOf(value))
		// fmt.Println("interface type", value.(type))

	}
}

func GenericOne() {
	var params []ITF = []ITF{
		4, 5, "atetet",
	}

	GenericExampleOne(params)

	var paramstwo []interface{} = []interface{}{4, 5, "23434"}
	fmt.Println("params", params, paramstwo)
}
