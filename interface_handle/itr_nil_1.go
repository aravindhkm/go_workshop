package interfacehandle

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

func ItrNilOne() {
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

// In golang, an interface is a type that specifies a set of method signatures. 
// When a value is assigned to an interface, golang constructs an interface value 
// that consists of two parts: the dynamic type and the dynamic value. This is 
// commonly referred to as the “interface tuple.”

// Dynamic Type: This is a pointer to a type descriptor that describes the type 
// of the concrete value stored in the interface.
// Dynamic Value: This is a pointer to the actual value that the interface holds.


// type iface struct {
//     tab  *itab
//     data unsafe.Pointer
// }

// type itab struct {
//     inter *interfacetype
//     _type *_type
//     hash  uint32
//     _     [4]byte
//     fun   [1]uintptr  // variable sized, actually [n]uintptr
// }

// tab: A pointer to an itab structure that contains information about the type and 
// the methods that the type implements for the interface.
// data: A pointer to the actual data held by the interface.
// When a value is assigned to an interface, golang finds the type descriptor for the 
// concrete type being assigned to the interface. Then sets up the method table (itab) 
// that allows method calls through the interface to be dispatched to the correct 
// implementation and finally stores a pointer to the actual value in the data field 
// of the interface.

// When aType = aImpl is executed

// Determining Interface Implementation: golang first determines that *SomeImpl (a pointer 
// to SomeImpl) implements the SomeType interface because *SomeImpl has a method Get() 
// with the correct signature.
// Looking Up the Type Descriptor: golang looks up the type descriptor for *SomeImpl.
// Creating the itab: golang creates an itab structure
// Assigning the Pointer: golang assigns the pointer to the SomeImpl value to the 
// data field of the interface.


// interface (aType)
// +------------+           +-----------+
// | tab        |---------->|  itab     |
// |            |           |-----------|
// | data       |--+        |  inter    |
// +------------+  |        |  _type    |
//                 |        |  fun[0]   |
//                 |        +-----------+
//                 |
//                 v
//          +---------------+
//          |   *SomeImpl   |
//          +---------------+
//          |   ........    |
//          +---------------+


// To sum up, what we learned, the previous explanation in short means:

// + uninitialised +-------------+ initialised +

// interface (aType)             interface (aType)
// +------------+                +--------------------------------+
// | tab:  nil  |                | tab:  type of *SomeImpl        |
// | data: nil  |                | data: value of *SomeImpl (nil) |
// +------------+                +--------------------------------+


// https://dev.to/crusty0gphr/tricky-golang-interview-questions-part-5-interface-nil-2agh