package sortingandsearching

import (
	"fmt"
)

// https://leetcode.com/explore/interview/card/top-interview-questions-medium/110/sorting-and-searching/802/


// Given an array of integers nums sorted in non-decreasing order, find the starting and 
// ending position of a given target value.
// If target is not found in the array, return [-1, -1].
// You must write an algorithm with O(log n) runtime complexity.

func searchRangeBinary(arr []int, target int) []int {
    indices := []int{}

    left, right := 0, len(arr)-1

    for left <= right {
        mid := left + (right-left)/2

		fmt.Println("mid", mid,"left", left, "right", right)
		// 5,7,7,8,8,10
        // Check if target is present at mid
        if arr[mid] == target {
            // Store the index of the target
            indices = append(indices, mid)

            // Check for other occurrences to the left
            i := mid - 1
            for i >= 0 && arr[i] == target {
                indices = append(indices, i)
                i--
            }

            // Check for other occurrences to the right
            i = mid + 1
            for i < len(arr) && arr[i] == target {
                indices = append(indices, i)
                i++
            }

            return indices
        }

        // If target is greater, ignore left half
        if arr[mid] < target {
            left = mid + 1
        } else { // If target is smaller, ignore right half
            right = mid - 1
        }
    }

    // If we reach here, then the element was not present
    return []int{-1}
}

func SortingAndSearchingExecuteBinary() {
	numsOne := []int{5,7,7,8,8,10}
	targetOne := 8
	result_one := searchRangeBinary(numsOne,targetOne)
	fmt.Println("resultOne", result_one)


	// nums_two := []int{5,7,7,8,8,10}
	// target_two := 6
	// result_two := searchRangeBinary(nums_two,target_two)
	// fmt.Println("result_two", result_two)


	// nums_three := []int{}
	// target_three := 0
	// result_three := searchRangeBinary(nums_three,target_three)
	// fmt.Println("result_three", result_three)
	
}

