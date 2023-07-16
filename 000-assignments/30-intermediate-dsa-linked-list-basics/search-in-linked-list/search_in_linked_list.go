package searchinlinkedlist

import (
	ll "linkedlist"
)

// Given a singly linked list and an integer B, search the linked list for a node that has value equal to B and return the node.
//
// Time Complexity: O(N)
// Space Complexity: O(1)
func SearchInLinkedList(A *ll.Node, B int) int {
	for A != nil {
		if A.Value == B {
			return 1
		}
		A = A.Next
	}
	return 0
}
