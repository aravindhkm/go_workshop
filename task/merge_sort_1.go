package task

import "fmt"

func reverseOwn(input []int) []int {
	mid := len(input) / 2
	for i, j := 0, len(input)-1; i <= mid && j >= mid; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
	return input
}

func reverse(input []int) []int {
	start := 0
	end := len(input) - 1

	for start < end {
		input[start], input[end] = input[end], input[start]
		start++
		end--
	}

	return input
}

func merge(left, right []int) ([]int, []int) {
	if len(left) == 0 || len(right) == 0 {
		return left, right
	}

	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] > left[j] {
			left[j] = left[i]
			j++
		} else {
			left[i] = left[j]
			i++
		}
	}

	return left, right
}

func mergeSort(input []int) {
	
}

func MergeSortOne() {
	input := []int{2, 1, 5, 3, 7, 8}
	reverseResult := reverse(input)
	fmt.Println("reverse", reverseResult)

	fmt.Println("merge sort")
}
