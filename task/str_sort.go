package task

import "fmt"

func customSorting(strArr []string) []string {
	oddStr := []string{}
	evenStr := []string{}
	result := []string{}
	for _, value := range strArr {
		if len(value)%2 != 0 {
			oddStr = append(oddStr, value)
		} else {
			evenStr = append(evenStr, value)
		}
	}
	fmt.Println("oddStr", oddStr)

	for i := 0; i < len(oddStr)-1; i++ {
		for j := 0; j < len(oddStr)-i-1; j++ {
			if len(oddStr[j]) > len(oddStr[j+1]) ||
				(len(oddStr[j]) == len(oddStr[j+1]) && oddStr[j] > oddStr[j+1]) {
				oddStr[j], oddStr[j+1] = oddStr[j+1], oddStr[j]
			}
		}
		fmt.Println(i, oddStr)

	}

	for i := 0; i < len(evenStr)-1; i++ {
		for j := 0; j < len(evenStr)-i-1; j++ {
			if len(evenStr[j]) < len(evenStr[j+1]) ||
				(len(evenStr[j]) == len(evenStr[j+1]) && evenStr[j] > evenStr[j+1]) {
				evenStr[j], evenStr[j+1] = evenStr[j+1], evenStr[j]
			}
		}
	}

	result = append(result, oddStr...)
	result = append(result, evenStr...)
	return result
}

func StrSort() {
	fmt.Println(customSorting([]string{"abc", "ab", "abcde", "a", "abcd"}))
	// output = [a abc abcde abcd ab]
}
