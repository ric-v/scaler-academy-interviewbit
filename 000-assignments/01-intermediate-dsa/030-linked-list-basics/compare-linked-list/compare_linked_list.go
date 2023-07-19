package comparelinkedlist

import (
	ll "linkedlist"
)

// CompareLinkedList compares two linked lists and returns 1 if they are equal and 0 if they are not.
// The function takes two linked lists as input and returns an integer.
// 
// Example: A : 1->2->3->4 and B : 1->2->3->4 are equal. 
// A : 1->2->3->4 and B : 1->2->3->5 are not equal.
// 
// Time Complexity: O(n)
// Space Complexity: O(1)
func CompareLinkedList(A, B *ll.Node) int {
	if A == nil && B == nil {
		return 1
	}
	if A == nil || B == nil {
		return 0
	}

	// Compare the values of the nodes and recursively call the function on the next nodes.
	if A.Value != B.Value {
		return 0
	}
	return CompareLinkedList(A.Next, B.Next)
}
