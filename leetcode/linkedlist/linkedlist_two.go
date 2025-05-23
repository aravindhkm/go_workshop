package linkedlist

import "fmt"

type NodeTwo struct {
	data     int
	nextNode *NodeTwo
}

func (n *NodeTwo) Add(data int) {
	if n.data == 0 {
		n.data = data
	} else {
		for {
			// fmt.Println(n.nextNode)
			if n.nextNode == nil {
				n.nextNode = &NodeTwo{data: data}
				break
			} else {
				n = n.nextNode
			}
		}
	}
}

func (n *NodeTwo) Print() {
	for {
		result := n.data
		if n.nextNode == nil {
			fmt.Println(result)
			break
		} else {
			result = n.data
			n = n.nextNode
		}

		fmt.Println(result)
	}
}
func LinkedListExecuteTwo() {
	node := NodeTwo{}

	node.Add(10)
	node.Add(20)
	node.Add(30)
	node.Add(40)

	node.Print()

	fmt.Println("LinkedListExecuteTwo")
}
