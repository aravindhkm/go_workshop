/*
Problem:
- Delete a node from a singly-linked list, given only a pointer to that node.

Approach:
- Since we don't have access to the previous node, simply copy the value and
  pointer of the next node and copy them into the current node.

Solution:
- Cache the next node.
- If the next node is nil, it's the last node. Just simply return.
- Copy the current node's value to the next node's value
- Copy the node's pointer to the next node's pointer.

Cost:
- O(1) time and O(1) space.
*/

package linkedlist

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func (d *Node) Add(new int) {
	if d.Data == 0 {
		d.Data = new
	} else {
		newNode := &Node{Data:new}

		if d.Next == nil {
			d.Next = newNode
		} else {
			currNode := d.Next

			for currNode.Next != nil {
				currNode = currNode.Next
			}

			currNode.Next = newNode
		}
	}
}

func (d *Node) List() {
	listNode := d
	for listNode != nil {
		fmt.Println(listNode.Data)

		if listNode.Next == nil {
			break
		}

		listNode = listNode.Next
	}
}


func (d *Node) Delete() {
	listNode := d
	for listNode != nil {
		fmt.Println(listNode.Data)

		if listNode.Next == nil {
			break
		}

		listNode = listNode.Next
	}
}

// deleteNode deletes the node in the linked list given only access to that node
func deleteNode(node *Node) {
    if node == nil || node.Next == nil {
        return
    }
    node.Data = node.Next.Data
    node.Next = node.Next.Next
}


// helper function to print the linked list
func printList(head *Node) {
    for head != nil {
        fmt.Printf("%d -> ", head.Data)
        head = head.Next
    }
    fmt.Println("nil")
}

func DeleteNodeMain() {
	var currNode Node

	currNode.Add(1)
	currNode.Add(2)
	currNode.Add(3)
	currNode.Add(4)
	currNode.Add(5)
	currNode.Add(6)

	currNode.List()

	head := &Node{Data: 1}
    head.Next = &Node{Data: 2}
    head.Next.Next = &Node{Data: 3}
    head.Next.Next.Next = &Node{Data: 4}

	fmt.Println("Original list:")
    printList(head)

	// Delete the node with value 3
	deleteNode(head.Next.Next)

	fmt.Println("List after deleting node with value 3:")
	printList(head)

}