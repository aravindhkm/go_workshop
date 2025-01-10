package sort

import (
	"fmt"
	"sort"
)

type data []string

func (v data) Len() int {
	return len(v)
}

func (v data) Less(i, j int) bool {
	return len(v[i]) < len(v[j])
}


func (v data) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func MapStringSort() {
	strSlice := []string{"banana", "apple", "orange",
	 "grape","Colorado", "Utah", "Wisconsin", "Oregon"}
	sort.Sort(data(strSlice))

	fmt.Println("Strings sorted by length:",strSlice)
	// for _, str := range strSlice {
	// 	fmt.Println(str)
	// }
}