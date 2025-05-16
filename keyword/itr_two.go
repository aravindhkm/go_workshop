package keyword

import (
	"fmt"
	"reflect"
)

type SomeType interface {
	Get() string
}

type SomeImpl struct {
	Name string
}

func (i SomeImpl) Get() string{
	i.Name = "world"
	return i.Name
}

func ItrTwo() {
	var aType SomeType
	if aType == nil {
		fmt.Println("nil interface", &aType)
	}

	var aImpl *SomeImpl
	if aImpl == nil {
		fmt.Println("nil struct", &aImpl)
	}

	aType = aImpl
	if aType == nil {
		fmt.Println("nil assignment", aType)
	} else {
		fmt.Println(reflect.TypeOf(aType))
		fmt.Println("Not nil assignment", &aType, aType)
	}

	aType = &SomeImpl{"Hello"}
	fmt.Println("aImpl   1", aImpl)
	fmt.Println("aType   2", aType.Get())

}

// In golang, an interface is a type that specifies a set of method signatures. When a value is assigned to an interface, golang constructs an interface value that consists of two parts: the dynamic type and the dynamic value. This is commonly referred to as the “interface tuple.”

// Dynamic Type: This is a pointer to a type descriptor that describes the type of the concrete value stored in the interface.
// Dynamic Value: This is a pointer to the actual value that the interface holds.
