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

func SecondLargetNumberOne() {
	input := []int{12, 5, 4, 13, 11, 21, 27}
	fmt.Println("Second Largest Number", secondLargetNumberFind(input))
}