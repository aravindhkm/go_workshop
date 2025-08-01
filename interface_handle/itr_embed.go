package interfacehandle

import "fmt"

// Basic interfaces
type A interface {
	Foo()
}

type B interface {
	Bar()
}

// Embedded interface
type C interface {
	A
	B
}

// Struct implementing all methods
type MyType struct{}

func (m MyType) Foo() {
	fmt.Println("Foo called")
}

func (m MyType) Bar() {
	fmt.Println("Bar called")
}

func useC(c C) {
	c.Foo()
	c.Bar()
}

func InterfaceEmbed() {
	var obj MyType
	useC(obj)
}
