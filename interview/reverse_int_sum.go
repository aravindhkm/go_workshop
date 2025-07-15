package interview

import "fmt"

func reverseAndSumNo(n int) (int, int) {
	reserved, totalSum := 0, 0

	for n != 0 {
		digit := n % 10
		reserved = reserved*10 + digit
		totalSum += digit
		n = n / 10
	}

	return reserved, totalSum
}

func ReverseIntAndSum() {
	n := 12347

	r1, r2 := reverseAndSumNo(n)
	fmt.Println("Result", r1, r2)
}

// | Metric           | Value                                                |
// | ---------------- | ---------------------------------------------------- |
// | Time Complexity  | **O(log₁₀(n))** — number of digits in `n`            |
// | Space Complexity | **O(1)** — constant space, no arrays or strings used |
