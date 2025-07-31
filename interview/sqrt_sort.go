package interview

import "fmt"

func sqrtSwap(data []int) []int {
	left, right := 0, len(data)-1
	pos := right
	result := make([]int, len(data))

	for left < right {
		l_data := data[left] * data[left]
		r_data := data[right] * data[right]

		if r_data > l_data {
			result[pos] = r_data
			right--
		} else {
			result[pos] = l_data
			left++
		}
		pos--
	}
	return result
}

// [-4, -1, 0, 3, 10]
// [16, 1, 0, 9, 100] // reverse pos
// [0, 1, 9, 16, 100]

func SqrtSort() {
	input := []int{-4, -1, 0, 3, 10}

	result := sqrtSwap(input)
	fmt.Println(result)
}

// Given an integer array nums sorted in non-decreasing order,
// return an array of the squares of each number sorted in non-decreasing order.
// Input: nums = [-4,-1,0,3,10]
// Output: [0,1,9,16,100]

// Two-Pointer Technique

// Since the array is sorted, the largest square values will be at the ends
// of the array (because of large negative or positive values).

// Use two pointers (left and right) to scan from both ends, compare their
// absolute values, and insert the square of the larger one at the end
// of the result array (from back to front).

// 🕒 Time Complexity: O(n)
// 📦 Space Complexity: O(n)
