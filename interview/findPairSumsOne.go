package interview

import "fmt"

func findPairSums(nums []int) map[int][][]int {
	pairSums := make(map[int][][]int)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			pairSums[sum] = append(pairSums[sum], []int{nums[i], nums[j]})

		}
	}

	return pairSums
}

func FindPairSumsOne() {
	nums := []int{1, 2, 3, 4, 2, 11, 5, 7, 8, 1}

	pairSums := findPairSums(nums)

	for i, pairSum := range pairSums {
		if len(pairSum) > 3 {
			fmt.Println(i, pairSum)
		}

	}
}