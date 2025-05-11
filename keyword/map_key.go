package keyword

import (
	"fmt"
)

func MapKey() {
	keyMap := make(map[int]string, 0)

	keyMap[1] = "hello"

	// var varKeyMap map[int]string

	// varKeyMap[2] = "hello 2"

	fmt.Println("keyMap", keyMap)
	// println("varKeyMap", varKeyMap)
}
