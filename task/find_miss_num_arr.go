package task

import (
	"fmt"
	"sort"
)


// var data sort.IntSlice

type unsort_arr []int

func (d unsort_arr) Len() int {
	return len(d)
}

func (d unsort_arr) Less(i, j int) bool {
	return d[i] > d[j]
}

func (d unsort_arr) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func findMissNumUnsortedArr(data []int) {
	sort.Sort(unsort_arr(data))
}

func FindMissingNumberInUnsortedArray() {
	miss_arr_one := []int{1, 6, 3, 2, 5, 8, 4} // expected result 7
	// miss_arr_one := []int{1, 6, 3, 2, 5}
	findMissNumUnsortedArr(miss_arr_one)
	fmt.Println("find_one", miss_arr_one)
}
