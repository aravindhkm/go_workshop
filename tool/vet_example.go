package tool

import "fmt"

func VetExample() {
	var x int // incorrect

	// var x int // correct one
	fmt.Printf("%s", x) // This is incorrect: %s expects a string, but x is an int
}

// go vet main.go
