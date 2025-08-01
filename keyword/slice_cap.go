package keyword

import "fmt"

// In Go, when you create a slice, it doesn’t actually hold the data itself —
// instead, it acts like a view or window into an array. That array —
// the one that holds the actual data — is called the underlying array.

// s[:0] → Called reslicing to zero length
// Not deletion or clearing — just trimming the view

func printSlice(s []int, cut int) {
	fmt.Printf("Full Slice: %v\n", s)
	if cut <= cap(s) {
		fmt.Printf("Slice[:%d]: %v\n", cut, s[:cut])
	} else {
		fmt.Printf("Slice[:%d]: out of capacity range\n", cut)
	}
	fmt.Printf("Length: %d, Capacity: %d\n", len(s), cap(s))
	fmt.Println("----------------------------------")
}

func SliceExampleOne() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s, 6)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s, 6)
	// fmt.Println("Checking", s[0]) // Error: index out of range

	// Extend its length.
	s = s[:4]
	printSlice(s, 6)
	// fmt.Println("Checking", s[5]) // Error: index out of range

	// Drop its first two values.
	s = s[2:]
	printSlice(s, 4)
	// fmt.Println("Checking", s[3]) // Error: index out of range if len < 4
}

func SliceAndCap() {
	SliceExampleOne()
}

// ==========================================================
// 🚀 Go Slice Memory Diagram - Step-by-Step
// ==========================================================

// 🧠 Underlying Array in Memory:
// Index:     0   1   2   3   4   5
// Value:    [2   3   5   7  11  13]
// Address:   ↑
//            Base pointer (initially points to index 0)

// ----------------------------------------------------------

// ✅ Step 1: Initial Slice
// Code: s := []int{2, 3, 5, 7, 11, 13}

// Slice `s`:
//   len = 6, cap = 6
//   View → [2   3   5   7  11  13]
//           ↑   ↑   ↑   ↑   ↑   ↑
// Base → points to index 0

// ----------------------------------------------------------

// ✅ Step 2: s = s[:0]

// Slice `s`:
//   len = 0, cap = 6
//   View → []
//   Underlying array remains:
//           [2   3   5   7  11  13]
//            ↑
//         Base → still index 0

// ----------------------------------------------------------

// ✅ Step 3: s = s[:4]

// Slice `s`:
//   len = 4, cap = 6
//   View → [2   3   5   7]
//           ↑   ↑   ↑   ↑
// Base → still index 0

// ----------------------------------------------------------

// ✅ Step 4: s = s[2:]

// Slice `s`:
//   len = 2, cap = 4
//   View → [5   7]
// Underlying array still:
//           [2   3   5   7  11  13]
//                    ↑   ↑
//                  Base   Base+1

// Base pointer has shifted to index 2

// ==========================================================
// 🧠 Key Points:

// - Slice is just a view: pointer + length + capacity
// - s[:0] hides elements but doesn't delete them
// - s[:4] shows previously hidden values
// - s[2:] moves base pointer forward
// - Underlying array remains intact unless reallocated
// ==========================================================
