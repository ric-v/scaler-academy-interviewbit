package insertinlinkedlist

import (
	"fmt"
	ll "linkedlist"
)

// InsertInLinkedList inserts a node with value B at position C in linked list A
// and returns the head of the linked list.
//
// Example: InsertInLinkedList(NewNode(1, 2), 3, 0) should return NewNode(3, 1, 2)
//
// Example: InsertInLinkedList(NewNode(1, 2), 3, 1) should return NewNode(1, 3, 2)
//
// Time Complexity: O(C)
// Space Complexity: O(1)
func InsertInLinkedList(A *ll.Node, B, C int) *ll.Node {
	var temp *ll.Node = A

	// If C is 0, insert at the beginning of the linked list
	if C == 0 {

		// Create a new node with value B and set its next to the head of the linked list
		head := ll.NewNode(B)
		head.Next = A
		return head
	} else {

		// Traverse the linked list until C is 0 and insert the node
		// with value B at that position and return the head of the linked list
		// If C is greater than the length of the linked list, insert at the end
		// of the linked list
		for temp.Next != nil {
			C--
			if C == 0 {
				newNode := ll.NewNode(B)
				newNode.Next = temp.Next
				temp.Next = newNode
				break
			}
			temp = temp.Next
		}
	}
	fmt.Println(temp, C)

	// If C is greater than the length of the linked list, insert at the end
	// of the linked list
	if C > 0 {
		temp.Next = ll.NewNode(B)
	}
	return A
}
