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
	fmt.Println("copyP2", copyP2)

	numPtr := new(int)
	fmt.Println("Value of numPtr before assignment:", *numPtr) // Output: 0 (zero-initialized)
	*numPtr = 42
	fmt.Println("Value of numPtr after assignment:", *numPtr) // Output: 42

	// Create a slice of int with length 3
	slice := make([]int, 3)
	slice[0] = 10
	slice[1] = 20
	slice[2] = 30
	fmt.Println("Slice:", slice)

	m := new(map[string]int)
	// (*m)["one"] = 1           // ❌ panic: assignment to entry in nil map
	*m = make(map[string]int) // initialize the actual map
	(*m)["one"] = 1

	// Create a map from string to int
	dictionary := make(map[string]int)
	dictionary["apple"] = 3
	dictionary["banana"] = 5
	fmt.Println("Map:", dictionary)
}

// | Feature       | `make`                                             | `new`                                             |
// | ------------- | -------------------------------------------------- | ------------------------------------------------- |
// | Purpose       | Initializes and allocates built-in reference types | Allocates memory for any type and returns pointer |
// | Return type   | **Initialized value** (not pointer)                | **Pointer** to zero value                         |
// | Used with     | Only for **slice, map, channel**                   | For **any type** (structs, arrays, basic types)   |
// | Memory alloc? | Yes                                                | Yes                                               |
// | Zero value?   | Not just zeroed — also initialized for use         | Only zeroed memory, **not initialized** for use   |
