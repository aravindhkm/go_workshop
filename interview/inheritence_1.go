package interview

import "fmt"

// Base struct
type Animal struct {
	Name string
}

func (a Animal) Speak() {
	fmt.Println(a.Name, "makes a sound")
}

// Derived-like struct
type Dog struct {
	Animal // Embedded struct (like inheritance)
	Breed  string
}

func (d Dog) Bark() {
	fmt.Println(d.Name, "barks! Woof woof!")
}

func InheritenceOne() {
	d := Dog{
		Animal: Animal{Name: "Tommy"},
		Breed:  "Labrador",
	}

	d.Speak() // inherited method
	d.Bark()  // Dog's own method
}

// | Feature           | Go                        | Java/C++         |
// | ----------------- | ------------------------  | ---------------- |
// | Inheritance       | ❌ Not directly supported | ✅ Supported      |
// | Struct embedding  | ✅ Yes                    | ❌ No             |
// | Polymorphism      | ✅ Via interfaces         | ✅ Via interfaces |
// | Method overriding | ✅ Possible               | ✅ Supported      |
