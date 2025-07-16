package generics

import "fmt"

func Comp[T comparable] (t1 T,t2 T) bool {
	return t1 == t2
}

func ComparableOne() {
	fmt.Println("Print", Comp(1, 1.1))
}


// example code 

// // Index returns the index of x in s, or -1 if not found.
// func Index[T comparable](s []T, x T) int {
// 	for i, v := range s {
// 		// v and x are type T, which has the comparable
// 		// constraint, so we can use == here.
// 		if v == x {
// 			return i
// 		}
// 	}
// 	return -1
// }

// func ExampleTwo() {
// 	// Index works on a slice of ints
// 	si := []int{10, 20, 15, -10}
// 	fmt.Println(Index(si, 15))

// 	// Index also works on a slice of strings
// 	ss := []string{"foo", "bar", "baz"}
// 	fmt.Println(Index(ss, "hello"))
// }
