package generics

import (
	"fmt"
	"reflect"
)

type ITF interface {}

func GenericExampleOne[T ITF] (data []T){
	for index, value := range data {
		fmt.Println("index", index)
		fmt.Println("index", reflect.TypeOf(value))

	}
}

func ExampleOne() {
	var params []ITF = []ITF{
		4,5,"atetet",
	}

	GenericExampleOne(params)

	// var params []interface{} = []interface{}{4,5,"23434"}
	fmt.Println("params",params)
}