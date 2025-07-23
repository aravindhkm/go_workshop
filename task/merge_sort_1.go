package task

import (
	"fmt"
)

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

func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			result = append(result, right[j])
			j++
		} else {
			result = append(result, left[j])
			i++
		}
	}

	return result
}

func mergeSort(input []int) []int {
	if len(input) <= 1 {
		return input
	}
	mid := len(input) / 2
	left := mergeSort(input[mid:])
	right := mergeSort(input[:mid])

	return merge(left, right)
}

func MergeSortOne() {
	input := []int{2, 1, 5, 3, 7, 8}
	reverseResult := reverse(input)
	fmt.Println("reverse", reverseResult)

	mergeResult := mergeSort(input)
	fmt.Println("mergeResult",mergeResult)

	fmt.Println("merge sort")
}
