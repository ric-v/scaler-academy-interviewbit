package reverselinkedlist

import (
	ll "linkedlist"
)

// ReverseList reverses a linked list
//
// Time complexity: O(n)
// Space complexity: O(1)
func ReverseList(A *ll.Node) *ll.Node {
	var prev, next *ll.Node
	curr := A

	// iterate through the list and reverse the pointers of each node until the end of the list
	for curr != nil {

		// save the next node and reverse the pointer of the current node
		next, curr.Next = curr.Next, prev

		// move to the next node in the list and update the previous node
		prev, curr = curr, next
	}

	return prev
}
