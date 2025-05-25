package keyword

import "fmt"

func Copy_A() {
	slice_a := []int{4, 5, 6, 7}
	arr_a := [4]int{6, 7, 8, 1}

	fmt.Println("slice_a", slice_a)
	fmt.Println("arr_a", arr_a)

	// copy
	// var result_slice []int // it not assign need to use make only
	// result_slice = make([]int, len(slice_a)+ 10)
	// copy(result_slice, slice_a)
	// fmt.Println("result_slice", result_slice, len(result_slice))

	// it won't work in array
	// result_slice_two := [4]int{}
	// copy(result_slice_two, slice_a)
	// fmt.Println("result_slice_two", result_slice_two, len(result_slice_two))

	// value := make([]string,0) // slice with initial length
	// value = append(value,"ak","ka")
	// fmt.Println("value",value, len(value))

	// map_value := make(map[int]string)
	// map_value[1] = "hello"
	// fmt.Println("map_value", map_value)

	// source := []int{1, 2, 3, 4, 5}
    // destination := make([]int, 3)
    
    // n := copy(destination, source)
    
    // fmt.Println("Number of elements copied:", n)
    // fmt.Println("Destination slice:", destination)

	fmt.Println("")
}
