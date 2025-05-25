package syncexample

import (
	"fmt"
	"sync/atomic"
)

func atomicBool() {
	atmoicBool := &atomic.Bool{}

	fmt.Println("before atmoicBool", atmoicBool.Load())
	atmoicBool.Store(true)
	fmt.Println("after atmoicBool", atmoicBool.Load())
	oldBool := atmoicBool.Swap(false)
	fmt.Println("oldBool", oldBool, atmoicBool.Load())
}

func atomicPtrInt() {
	var atomicPtr atomic.Pointer[int]

	val1 := 42
	atomicPtr.Store(&val1)

	ptr := atomicPtr.Load()
	fmt.Println("Loaded value:", *ptr) // Output: Loaded value: 42

	val2 := 100
	oldPtr := atomicPtr.Swap(&val2)
	fmt.Println("Swapped old value:", *oldPtr) // Output: Swapped old value: 42

	ptr = atomicPtr.Load()
	fmt.Println("Current value:", *ptr) // Output: Current value: 100
}

type User struct {
	name    string
	contact int
}

func atomicPtrStruct() {
	usr := &User{"Hello", 1}

	var atomicPtr atomic.Pointer[User]

	atomicPtr.Store(usr)

	ptr := atomicPtr.Load()
	fmt.Println("Loaded value:", *ptr) // Output: Loaded value: 42

	usr2 := &User{"Ok", 2}
	oldPtr := atomicPtr.Swap(usr2)
	fmt.Println("Swapped old value:", *oldPtr) // Output: Swapped old value: 42

	ptr = atomicPtr.Load()
	fmt.Println("Current value:", *ptr) // Output: Current value: 100
}

type mapType map[string]int

func atomicPtrMap() {
	usrMap := &mapType{
		"Hello": 1,
	}

	var atomicPtr atomic.Pointer[mapType]

	atomicPtr.Store(usrMap)

	ptr := atomicPtr.Load()
	fmt.Println("Loaded value:", *ptr) // Output: Loaded value: 42

	usrMap2 := &mapType{
		"Ok": 2,
	}

	oldPtr := atomicPtr.Swap(usrMap2)
	fmt.Println("Swapped old value:", *oldPtr) // Output: Swapped old value: 42

	ptr = atomicPtr.Load()
	fmt.Println("Current value:", *ptr) // Output: Current value: 100
}

func Atomic_C() {
	// atomicBool()
	// atomicPtrInt()
	// atomicPtrStruct()
	atomicPtrMap()
}
