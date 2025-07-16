package interview

import "fmt"

func PalinGet(x int) bool {
	// Negative numbers and numbers ending in 0 (but not 0 itself) are not palindromes
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversed := 0
	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}

	// For odd digit numbers, ignore the middle digit by reversed/10
	return x == reversed || x == reversed/10
}

func PalinOwn(x int) bool {
	lastDigit := 0
	for x > lastDigit {
		if lastDigit == 0 {
			lastDigit = x % 10
		}
		// fmt.Println("before cut", x)
		// cutLastDigit
		x /= 10
	}
	// fmt.Println(lastDigit, x)
	return x == lastDigit
}

func Palindrome_One() {

	testcase := []int{121, -121, 135, 272, 10, 12321, 0, 1221, 123} //, 135, 272, 10, 12321, 0, 1221, 123

	for _, value := range testcase {
		fmt.Println(value, PalinGet(value))
	}
}


// Time complexity: O(log₁₀(n))

// Space complexity: O(1)