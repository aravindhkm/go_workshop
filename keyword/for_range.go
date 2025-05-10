package keyword

import "fmt"

func ForRange() {
	arr := [3]int{3,4,5}

	for index, value := range arr {
		fmt.Println("Array value", index, value)
	}

	fmt.Println("")

	slice := []int{3,4,5}
	for index, value := range slice {
		fmt.Println("Slice value", index, value)
	}

	fmt.Println("")

	keyMap := make(map[int]string, 4)
	keyMap[3] = "Hello 1"
	keyMap[4] = "Hello 2"
	keyMap[5] = "Hello 3"

	for key, value := range keyMap {
		fmt.Println("KeyMap value", key, value)
	}

}