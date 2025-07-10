package interview

import (
	"fmt"
)

// Node represents a node in the singly linked list.
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// LinkedList represents the singly linked list.
type LinkedList[T comparable] struct {
	head *Node[T]
	size int
}

// InsertAtHead inserts a new node at the beginning.
func (ll *LinkedList[T]) InsertAtHead(value T) {
	newNode := &Node[T]{Value: value, Next: ll.head}
	ll.head = newNode
	ll.size++
}

// InsertAtTail inserts a new node at the end.
func (ll *LinkedList[T]) InsertAtTail(value T) {
	newNode := &Node[T]{Value: value}
	if ll.head == nil {
		ll.head = newNode
	} else {
		curr := ll.head
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = newNode
	}
	ll.size++
}

// Delete deletes the first node with the given value.
func (ll *LinkedList[T]) Delete(value T) bool {
	if ll.head == nil {
		return false
	}
	if ll.head.Value == value {
		ll.head = ll.head.Next
		ll.size--
		return true
	}

	prev := ll.head
	curr := ll.head.Next
	for curr != nil {
		if curr.Value == value {
			prev.Next = curr.Next
			ll.size--
			return true
		}
		prev = curr
		curr = curr.Next
	}
	return false
}

// Search checks if a value exists in the list.
func (ll *LinkedList[T]) Search(value T) bool {
	curr := ll.head
	for curr != nil {
		if curr.Value == value {
			return true
		}
		curr = curr.Next
	}
	return false
}

// Reverse reverses the linked list in-place.
func (ll *LinkedList[T]) Reverse() {
	var prev *Node[T]
	curr := ll.head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	ll.head = prev
}

// Size returns the number of elements in the list.
func (ll *LinkedList[T]) Size() int {
	return ll.size
}

// Display prints the linked list.
func (ll *LinkedList[T]) Display() {
	curr := ll.head
	for curr != nil {
		fmt.Printf("%v -> ", curr.Value)
		curr = curr.Next
	}
	fmt.Println("nil")
}

// --- Example usage ---
func LinkedListExec() {
	ll := LinkedList[int]{}

	ll.InsertAtHead(10)
	ll.InsertAtHead(20)
	ll.InsertAtTail(5)
	ll.Display() // 20 -> 10 -> 5 -> nil

	ll.Delete(10)
	ll.Display() // 20 -> 5 -> nil

	fmt.Println("Search 5:", ll.Search(5)) // true
	fmt.Println("Size:", ll.Size())        // 2

	ll.Reverse()
	ll.Display() // 5 -> 20 -> nil
}
 