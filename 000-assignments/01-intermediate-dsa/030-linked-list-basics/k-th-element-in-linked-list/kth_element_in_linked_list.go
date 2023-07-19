package kthelementinlinkedlist

import (
	ll "linkedlist"
)

// KthElementInLinkedList returns the kth element in the linked list A
// 1 <= k <= length of the list A
// 1 <= length of the list A <= 10^5
//
// Example: 1 -> 2 -> 3 | B = 2 | return 3
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func KthElementInLinkedList(A *ll.Node, B int) int {
	var temp *ll.Node = A
	var count int

	// until the end of linked list
	for temp != nil {

		// if count is equal to B then return the value
		if count == B {
			return temp.Value
		}

		// increment count and move to next node
		count++
		temp = temp.Next
	}
	return 0
}
