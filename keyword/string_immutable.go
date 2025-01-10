package keyword

import "fmt"

func exampleOne() {
	str := "Hello, World!"
	fmt.Println(str)

	// Attempt to change the string
	// This won't work, as strings are immutable in Go
	// str[0] = 'h' // This would cause a compilation error

	// Instead, create a new string
	newStr := "h" + str[1:]
	fmt.Println(newStr)
}

func exampleTwo() {
	str := "Hello, World!"
	fmt.Println(str)

	// Convert string to []byte
	bytes := []byte(str)

	// Modify the byte slice
	bytes[0] = 'h'

	// Convert back to string
	newStr := string(bytes)
	fmt.Println(newStr)
}

func StringImmutable() {
	exampleOne()
	exampleTwo()
}