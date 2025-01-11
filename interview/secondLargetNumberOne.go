package interview

import "fmt"

func secondLargetNumberFind(nums []int) int {
	if len(nums) < 2 {
		return -1
	}
	largest, secondLargest := -1, -1

	for _, num := range nums {
		if num > largest {
			secondLargest = largest
			largest = num
		}
	}

	return secondLargest
}

func secondApporach(input []int) int{
	result := make([]int,2)

	for _ , value := range input {
		if result[0] < value {
			result[1] =  result[0]
			result[0] = value
		} else if result[1] < value {
			result[1] = value
		}
	}
	return result[1]
}

func SecondLargetNumberOne() {
	input := []int{12, 5, 4, 13, 11, 21, 27}
	fmt.Println("Second Largest Number", secondApporach(input))
}