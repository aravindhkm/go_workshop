package linkedlist

import "fmt"


type Node struct {
	Data int
	Next *Node
}

func (s *Node) Add(v int) {
	if s.Data == 0 {
		s.Data = v
	} else {
		current := s

		for current.Next != nil {
			current = current.Next

		}
		current.Next = &Node{Data: v}
	}
}

func (s *Node) Print() {
	for s != nil {
		fmt.Println("Data", s.Data)
		s = s.Next
	}
}

func LinkedListExecuteOne() {
	var data *Node

	data.Add(1)
	data.Add(7)
	data.Add(2)
	data.Add(11)
	data.Add(22)
	data.Add(44)

	data.Print()

	fmt.Println("Linked List Printed")

}