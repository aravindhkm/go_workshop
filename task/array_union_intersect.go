package task


import (
	"fmt"
)

func UnionIntersectExecute() {
	arr_one := []int{1,3, 4, 5,7}
	arr_two := []int{2, 3, 5, 6}

	union_result := []int{}
	intersect_result := []int{}
	union_map := make(map[int]bool)

	for _,value := range arr_one {
		if !union_map[value] {
			union_result = append(union_result, value)
			union_map[value] = true
		}
	}

	for _,value_arr_two := range arr_two {
		if !union_map[value_arr_two] {
			union_result = append(union_result, value_arr_two)
			union_map[value_arr_two] = true
		}
	}

	for _, value := range arr_one {
		for _, value_two := range arr_two {
			if value == value_two {
				intersect_result = append(intersect_result, value_two)
			}
		}
	}

	fmt.Println("union_result", union_result)
	fmt.Println("union_result", intersect_result)

	fmt.Println("")
}