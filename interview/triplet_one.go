package interview

import (
	"fmt"
	"sort"
)

func mergeTripletData(a, b []int) []int {
	left, right := 0, 0
	result := []int{}
	for left < len(a) && right < len(b) {
		if a[left] < b[right] {
			result = append(result, a[left])
			left++
		} else {
			result = append(result, b[right])
			right++
		}
	}

	result = append(result, a[left:]...)
	result = append(result, b[right:]...)

	return result
}

func mergeSortTriplet(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSortTriplet(arr[:mid])
	right := mergeSortTriplet(arr[mid:])

	return mergeTripletData(left, right)
}

func findTripletOne(arr []int, k int) [][]int {
	n := len(arr) - 1

	fmt.Println("merge", mergeSortTriplet(arr))

	sort.Ints(arr)
	fmt.Println("before start", arr)
	result := [][]int{}
	for i := 0; i < n-1; i++ {
		left := i + 1
		right := n - 2

		for left < right {
			sum := arr[i] + arr[left] + arr[right]

			if sum == k {
				result = append(result, []int{arr[i], arr[left], arr[right]})
				break
			} else if sum < k {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

func TripletOne() {
	result := findTripletOne([]int{47, 6, 3, 8, 12, 10, 5, 13, 21}, 28)
	fmt.Println("result", result)
}

// Find a triplet in [47, 6, 3, 8, 12, 10, 5, 13, 21]
// that will sums to k=28

// ✅ Best DSA Approach: Two Pointer Technique
// Time Complexity: O(n^2)

// Space Complexity: O(1) (if done in-place)

// 💡 Steps:
// Sort the array.

// Fix one element i, then use two pointers (left and right) to find the other two such that their sum is k - arr[i].

// Original array: [47, 6, 3, 8, 12, 10, 5, 13, 21]
// Sorted array:   [3, 5, 6, 8, 10, 12, 13, 21, 47]
//                              ↑    ↑         ↑
//                             i   left      right

// For i = 0 (3):
//     - left = 1 (5), right = 8 (47)
//     - sum = 3 + 5 + 47 = 55 → too high → move right--
//     - ...
//     - sum = 3 + 10 + 13 = 26 → too low → move left++
//     - sum = 3 + 12 + 13 = 28 ✅ Triplet found!

// Merge Sort on [38, 27, 43, 3, 9, 82, 10]

//                      [38, 27, 43, 3, 9, 82, 10]
//                       /                      \
//           [38, 27, 43]                        [3, 9, 82, 10]
//            /       \                           /          \
//      [38]       [27, 43]                  [3, 9]         [82, 10]
//                 /     \                   /    \          /    \
//              [27]   [43]              [3]     [9]     [82]    [10]

//                ↓ Merge steps (bottom-up) ↓

//            [27, 43]      → [27, 38, 43]
//       [3, 9] → [3, 9]      [82, 10] → [10, 82]

//                ↓ Final Merge ↓

//       [27, 38, 43] + [3, 9, 10, 82] → [3, 9, 10, 27, 38, 43, 82]
