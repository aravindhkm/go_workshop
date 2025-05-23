package task

import (
	"errors"
	"fmt"
)

type Queue struct {
	items []int
}

func (q *Queue) Push(newItem int) {
	q.items = append(q.items, newItem)
}

func (q *Queue) Pop() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Queue is Empty")
	}
	firstValue := q.items[0]

	newQueue := q.items[1:]
	q.items = newQueue

	return firstValue, nil
}

func (q *Queue) Print() {
	fmt.Println(q.items)
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}

func QueueOne() {
	var q Queue

	q.Push(10)
	q.Push(20)
	q.Push(30)

	fmt.Println("Queue Size", q.Size())
	q.Print()
	q.Pop()

	q.Push(40)
	q.Push(50)
	q.Print()
}
