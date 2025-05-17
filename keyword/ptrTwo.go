package keyword

import "fmt"

// slot 2
// holding value slot 0 memory address
func printPointer(param *string) {
	n := &param
	k := n
	fmt.Println("0", param)
	fmt.Println("result 1", *param)

	*param = "hello"
	fmt.Println("Ref", param)
	fmt.Println("result 2", *param)

	ak := "ak"
	*n = &ak

	fmt.Println("result 3", *param)

	fmt.Println("param addr ->", &param)
	fmt.Println("n addr ->", &n)
	fmt.Println("k addr ->", &k)
}

func prtStr(val *string) *string {
	return val
}

func printPointerVal(namePointer *string) {
	*namePointer = "hello"
	fmt.Println(*namePointer)
}

func PtrTwo() {
	// slot 0 = memory addr
	// holding value string
	name := "bill"

	// slot 1 = memory addr
	// holing value slot 0
	namePointer := &name

	fmt.Println(&name)
	fmt.Println(&namePointer)
	printPointer(namePointer)

	fmt.Println("final name", name)

	fmt.Println(name)
	fmt.Println(*namePointer)
	printPointerVal(namePointer)
	fmt.Println(name)
	fmt.Println(*namePointer)
}

// Question 6:
// Here's a tough one.

// We've been discussing about how to use pointers to avoid passing something to a function by value.  So instead of passing a value to a function, we pass a pointer to that value.  But we've also said many times that Go is a "pass by value" language, which *always* copies arguments that are passed to a function.  Take a glance at the following code snippet, which gets a pointer to name called namePointer and prints out the memory address of the pointer itself!

// Then in the function, we take the pointer that was passed as an argument and print out the address of that pointer too.

// Do you think the memory address printed by both Println calls will be the same?  Why or why not?

// package main

// import "fmt"

// func main() {
//  name := "bill"

//  namePointer := &name

//  fmt.Println(&namePointer)
//  printPointer(namePointer)
// }

// func printPointer(namePointer *string) {
//  fmt.Println(&namePointer)
// }
