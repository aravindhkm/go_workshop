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

	var m *MyStruct = nil
	doSomething(m)
}
