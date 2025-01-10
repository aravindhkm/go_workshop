package sortingandsearching

import (
	"fmt"
)

// https://leetcode.com/explore/interview/card/top-interview-questions-medium/110/sorting-and-searching/802/


// Given an array of integers nums sorted in non-decreasing order, find the starting and 
// ending position of a given target value.
// If target is not found in the array, return [-1, -1].
// You must write an algorithm with O(log n) runtime complexity.

func searchRange(nums []int, target int) []int {
    start := 0
	end := len(nums) - 1

	var result []int

	for start <= end {
		mid := end / 2

		isLoop := end - mid <= 5
		
		if target > nums[mid] {
			if isLoop {
				for i:=mid;i<=end;i++ {
					if nums[i] == target {
						result = append(result, i)
					}
				}

				break
			} 
		} else {
			if isLoop {
				for i:=0;i<=mid;i++ {
					if nums[i] == target {
						result = append(result, i)
					}
				}
				
				break
			}
		}
	}

	if len(result) > 0 {
		return result
	}

	result = append(result, -1,-1)

	return result
}

func SortingAndSearchingExecuteOwn() {
	numsOne := []int{5,7,7,8,8,10}
	targetOne := 8
	result_one := searchRange(numsOne,targetOne)
	fmt.Println("resultOne", result_one)


	nums_two := []int{5,7,7,8,8,10}
	target_two := 6
	result_two := searchRange(nums_two,target_two)
	fmt.Println("result_two", result_two)


	nums_three := []int{}
	target_three := 0
	result_three := searchRange(nums_three,target_three)
	fmt.Println("result_three", result_three)
	
}

