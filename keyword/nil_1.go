package keyword

import "fmt"

type ListNode struct {
	Data int
	Next *ListNode
}

func NilOne() {
	node1 := &ListNode{}
	// &ListNode{} creates a pointer to a newly allocated ListNode struct.
	// Even though all fields inside it are zero-values (Data = 0, Next = nil),
	// the pointer itself (node1) points to a valid memory location — so it’s not nil.

	var node2 *ListNode
	// You declare a variable of type *ListNode (pointer to ListNode) but don’t assign it anything.
	// By default, an uninitialized pointer in Go is nil.
	// So node2 literally holds the nil value (no memory address).

	var node3 *ListNode = &ListNode{}

	fmt.Println(node1 == nil) // false
	fmt.Println(node2 == nil) // true
	fmt.Println(node3 == nil) // false

	fmt.Println("Nil checker")
}
