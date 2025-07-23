package workout

import (
	"fmt"
)

type Node struct {
	Data     int
	NextNode *Node
}

func (d *Node) Add(newVal int) {
	if d.Data == 0 {
		d.Data = newVal
	} else {
		for {
			if d.NextNode == nil {
				d.NextNode = &Node{Data: newVal}
				break
			} else {
				d = d.NextNode
			}
		}
	}
}

func (d *Node) Get() {
	for {
		fmt.Println(d.Data)

		if d.NextNode != nil {
			d = d.NextNode
		} else {
			break
		}
	}
}

func DummyWorkOut() {
	// linkedList := Node{Data: 1}

	// linkedList.Add(2)
	// linkedList.Add(3)
	// linkedList.Add(4)

	// linkedList.Get()

	fmt.Println()
}
