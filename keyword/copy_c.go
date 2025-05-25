package keyword

import "fmt"

func removeAtIndex(s []int, index int) []int {
	if index < 0 || index >= len(s) {
		return s // invalid index
	}
	fmt.Println("1", s[index:])   // 30, 40, 50
	fmt.Println("2", s[index+1:]) // 40, 50

	copy(s[index:], s[index+1:])

	fmt.Println("s", s)
	return s[:len(s)-1]
}

func Copy_C() {
	s := []int{10, 20, 30, 40, 50}
	s = removeAtIndex(s, 2)
	fmt.Println(s) // Output: [10 20 40 50]

	// Create a large slice
	large := make([]byte, 1_000_000)
	small := large[:100] // small slice refers to the same large array

	// Memory leak risk: `small` still holds reference to entire large array

	// Correct way: copy to a new slice
	newSmall := make([]byte, 100)
	copy(newSmall, small)

	// Now large can be garbage collected
	large = nil // help GC
	fmt.Printf("newSmall length: %d, capacity: %d\n", len(newSmall), cap(newSmall))
}
