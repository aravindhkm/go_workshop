package interview

import "fmt"

func PalinGet(n int) bool {

	return false
}

func Palindrome_One() {

	testcase := []int{121, 135, 272}

	for _, value := range testcase {
		fmt.Println("Palindrome", PalinGet(value))
	}
}
