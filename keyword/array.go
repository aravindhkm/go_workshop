package keyword

import "fmt"

func ArrayTest() {
	var array_one [4]int
	array_one[0] = 1
	array_one[1] = 1
	array_one[2] = 1
	array_one[3] = 1
	fmt.Println("array_one", array_one)

	var array_two = [3]string{"hello", "world", "k"}
	fmt.Println("array_two", array_two)

	array_three := [4]string{"world", "k", "hello","fuck"}
	fmt.Println("array_three", array_three)

	array_four := make([]int,3,4)
	fmt.Println("array_four", array_four)
}