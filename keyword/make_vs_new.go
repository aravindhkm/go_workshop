package keyword

import "fmt"

type Person struct {
	Name string
}

func MakeVsNewExampleOne() {
	p1 := new(Person)
	p1.Name = "Alice"

	copyP1 := p1
	copyP1.Name = "Bob"
	
	p2 := Person{"Aravind"} 
	p2.Name = "Alice"

	copyP2 := p2
	copyP2.Name = "Bob"

	fmt.Println("p1", p1)
	fmt.Println("p2", p2)


	numPtr := new(int)
	fmt.Println("Value of numPtr before assignment:", *numPtr) // Output: 0 (zero-initialized)
	*numPtr = 42
	fmt.Println("Value of numPtr after assignment:", *numPtr)  // Output: 42

	// Create a slice of int with length 3
	slice := make([]int, 3)
	slice[0] = 10
	slice[1] = 20
	slice[2] = 30
	fmt.Println("Slice:", slice)

	// Create a map from string to int
	dictionary := make(map[string]int)
	dictionary["apple"] = 3
	dictionary["banana"] = 5
	fmt.Println("Map:", dictionary)
}
