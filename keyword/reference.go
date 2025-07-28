package keyword

import "fmt"

func sliceAppend(s []int, newVal int) {
	s = append(s, newVal)
}

func sliceIndexChange(s []int, index, value int) {
	s[index] = value
}

// func 

func ReferenceOne() {
	arr := []int{4, 5, 6}
	sliceAppend(arr, 8)

	fmt.Println(arr)

	sliceIndexChange(arr, 1, 10)
	fmt.Println(arr)

// 	mapData := make(map[string]int)

// 	mapData
}

//  ./run.sh keyword/ 16
