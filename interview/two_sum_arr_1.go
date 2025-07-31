package interview

import "fmt"

func twoSumHashMap(nums []int, target int) []int {
	keyMap := make(map[int]int)

	for index, value := range nums {
		hashValue := target - value
		if pair, ok := keyMap[hashValue]; ok {
			return []int{pair, index}
		}
		keyMap[value] = index
	}

	return nil
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int) // value -> index

	for i, num := range nums {
		complement := target - num
		if index, found := numMap[complement]; found {
			return []int{index, i}
		}
		numMap[num] = i
	}
	return nil // or panic("No two sum solution")
}

func TwoSumArrOne() {
	nums := []int{3, 7, 2, 11, 15, 2}
	target := 14

	fmt.Println(twoSum(nums, target))
}

// Input: nums = [2, 7, 11, 15], target = 9
// Output: [0, 1]  // Because nums[0] + nums[1] = 2 + 7 = 9

// ✅ Best Approach: Hash Map
// Time Complexity: O(n)
// Space Complexity: O(n)
// 👉 Idea:
// Traverse the array.

// Mention edge cases like:

// Duplicate numbers

// Negative numbers

// Large array sizes
