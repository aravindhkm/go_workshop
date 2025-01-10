package keyword

import "fmt"

func modifyArray(arr [3]int) {
	arr[0] = 100 // Modifying the copy
}

func modifySlice(s []int) {
	s[0] = 100 // Modifying the original slice
}

func modifyMap(m map[string]int) {
	m["key"] = 100 // Modifying the original map
}


func ReferenceOrValue() {
	// Array example
	arr := [3]int{1, 2, 3}
	modifyArray(arr)
	fmt.Println("Array after modification:", arr) // Output: [1 2 3] (unchanged)

	// Slice example
	slice := []int{1, 2, 3}
	modifySlice(slice)
	fmt.Println("Slice after modification:", slice) // Output: [100 2 3] (modified)

	originalMap := map[string]int{"key": 10, "anotherKey": 20}
	fmt.Println("Original map before modification:", originalMap)

	modifyMap(originalMap)
	fmt.Println("Original map after modification:", originalMap)
}