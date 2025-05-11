package keyword

import "fmt"

func StringImmutableOne() {
	str := "Hello, World!"
	fmt.Println(str)

	// Attempt to change the string
	// This won't work, as strings are immutable in Go
	// str[0] = 'h' // This would cause a compilation error

	// Instead, create a new string
	newStr := "h" + str[1:]
	fmt.Println(newStr)
	fmt.Println(" \n")
}

func StringImmutableTwo() {
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

// Why use rune instead of byte?
// byte = uint8 = 1 byte — good for raw data, ASCII, or byte slices.
// rune = int32 = 4 bytes — represents a full Unicode code point (character).
// Take the string "世界" (Chinese for "world"). It's 2 characters, but:
// s := "世界"
// fmt.Println(len(s))         // Outputs: 6 (each Chinese character = 3 bytes)
// fmt.Println(len([]rune(s))) // Outputs: 2
// StringReverseTwo is a bad approach

func StringReverseOne() {
	s := "Hello, 世界"

	runes := []rune(s) // Convert string to rune slice to handle Unicode
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // Swap runes
	}

	fmt.Println("Reversed One:", string(runes))
}

func StringReverseTwo() {
	str := "Hello, 世界"
	length := len(str)
	resultStr := make([]byte, length)

	for i := 0; i < length; i++ {
		resultStr[i] = str[length-1-i]
	}

	fmt.Println("Reversed Two:", string(resultStr))
}

func StringImmutable() {
	s := "世界"
	fmt.Println(len(s))         // Outputs: 6 (each Chinese character = 3 bytes)
	fmt.Println(len([]rune(s))) // Outputs: 2


	// StringImmutableOne()
	// StringImmutableTwo()
	StringReverseOne()
	// StringReverseTwo()
}
