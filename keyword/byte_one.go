package keyword

import "fmt"

// bytes is a pass by reference

func UpdateByte(Arg []byte) {
	copy(Arg, []byte("Test"))
	// 	Arg = []byte("Test")
}

func ByteOne() {
	byteArg := make([]byte, 0, 524)

	fmt.Println("First One", string(byteArg))

	byteArg = []byte("Hello")

	fmt.Println("Second One", string(byteArg))

	UpdateByte(byteArg)

	fmt.Println("Third One", string(byteArg))
}
