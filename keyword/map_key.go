package keyword

import (
	"fmt"
)

func MapKey() {
	keyMap := make(map[int]string)
	keyMap[1] = "hello"
	fmt.Println("keyMap:", keyMap)

	// var varKeyMap map[int]string declares a map variable, but does not initialize it.
	// This map is nil (has no memory allocated).
	// it will not work
	var varKeyMap map[int]string
	// it will works
	// var varKeyMap map[int]string = map[int]string{
	// 	3: "ge",
	// }
	varKeyMap[2] = "hello from variable"
	println("varKeyMap:", varKeyMap[2])
}
