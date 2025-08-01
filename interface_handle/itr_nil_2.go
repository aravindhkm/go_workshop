package interfacehandle

import "fmt"

type MyStruct struct{}

func doSomething(i interface{}) {
	if i == nil {
		fmt.Println("Received nil interface")
	} else {
		fmt.Println("Interface is not nil!") // This will print
	}
}

func ItrNilTwo() {
	doSomething(nil) // Output: Received nil interface

	var d []int
	fmt.Println("Check Nil ---->", d == nil)
	doSomething(d) // Output: Interface is not nil!

	var m *MyStruct = nil
	fmt.Println("Check Nil ---->", m == nil)
	doSomething(m) // Output: Interface is not nil!
	// If you pass a typed nil (e.g., a *MyStruct that is nil),
	// then the interface itself is not nil, because it holds a type.

}
