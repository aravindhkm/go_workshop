package task

import (
	"fmt"
)

func findChar(name string, charArr []rune) []bool {
	// Build a map of characters in name for O(1) lookup
	charMap := make(map[rune]bool)
	for _, c := range name {
		charMap[c] = true
	}

	// Check each character in charArr
	results := make([]bool, len(charArr))
	for i, r := range charArr {
		_, exists := charMap[r]
		results[i] = exists
	}

	return results
}

func FindChar() {
	name := "Hello"
	char := []rune{'q', 'o', 'a', 'e', 'l', 'q', 'r', 'H'}

	results := findChar(name, char)

	for _, result := range results {
		fmt.Println(result)
	}
}


// Algorithm - Hash Map Lookup
// Time Complexity - O(n + m)
// DSA Principle - Followed (hashing, clean logic)
