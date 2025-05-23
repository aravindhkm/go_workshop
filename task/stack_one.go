package task

import (
	"errors"
	"fmt"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	lastIndex := len(s.items) - 1
	top := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return top, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Print() {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Println("Stack (bottom -> top):")
	for i, val := range s.items {
		fmt.Printf("[%d] %d\n", i, val)
	}
}

func StackOne() {
	var stack Stack

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	stack.Print()

	fmt.Println("Stack size:", stack.Size())

	top, _ := stack.Peek()
	fmt.Println("Top element:", top)

	// for !stack.IsEmpty() {
	// 	val, _ := stack.Pop()
	// 	fmt.Println("Popped:", val)
	// }

	_, err := stack.Pop()
	if err != nil {
		fmt.Println("Error:", err)
	}

	stack.Print()

}
