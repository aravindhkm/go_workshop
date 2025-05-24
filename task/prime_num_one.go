package task

import (
	"fmt"
	"math"
)

func isPrime(data int) bool {
	if data <1 {
		return true
	}

	sqrt := int(math.Sqrt(float64(data)))

	for i:=2; i<=sqrt; i++ {
		if data % i == 0 {
			return false
		}
	}
	return true
}

func PrimeNumOne() {
	for i := 1; i < 12; i++ {
		if isPrime(i) {
			fmt.Println(i, "is a prime number ---------->")
		} else {
			fmt.Println(i, "is not a prime number")

		}
	}
}

// Name: Trial Division Algorithm

// Algorithm: Trial Division
// Type: Basic Primality Testing
// DSA Category: Number Theory / Brute Force Search

// Time Complexity: O(√n)

// Because we only check up to √n divisors.

// Space Complexity: O(1)

// No extra space used beyond a few variables.
