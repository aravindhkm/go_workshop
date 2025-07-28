package keyword

import "fmt"

func SliceTest() {
	// single slice
	var array_one []int
	array_one = append(array_one, 4,5,6)
	fmt.Println("array_one", array_one)

	var array_two = []int{1,2,3,4}
	fmt.Println("array_two", array_two)

	array_three := []string{}
	array_three = append(array_three, "a", "b", "c")
	fmt.Println("array_three", array_three)

	// two dimensional slice
	var two_array_one [][]int
	two_array_one = append(two_array_one, []int{1,5}, []int{11,55,6})
	fmt.Println("two_array_one",two_array_one)

	var two_array_two = [][]int{{5,6},{12,546,6,7,8,7}}
	fmt.Println("two_array_two",two_array_two)

	two_array_three := [][]int{{5,6},{12,546,6,7,8,7}}
	fmt.Println("two_array_three",two_array_three)

	// three dimensional slice
	var three_array_one [][][]int
	three_array_one = append(three_array_one, [][]int{{1,5}}, [][]int{{11,55,6}})
	fmt.Println("three_array_one",three_array_one)

	var three_array_two = [][][]int{{{5,6,7},{33,44}}}
	fmt.Println("three_array_two",three_array_two)
}