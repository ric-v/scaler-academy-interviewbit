package middleelementoflinkedlist

import (
	ll "linkedlist"
)

// MiddleElementOfLinkedList returns the middle element of a linked list
// Time complexity: O(n)
// Space complexity: O(1)
func MiddleElementOfLinkedList(A *ll.Node) (int, *ll.Node) {

	// If the linked list is empty, return 0
	if A == nil {
		return 0, nil
	}

	var slow, fast *ll.Node = A, A

	// If the linked list has odd number of elements, fast.Next.Next will be nil
	// If the linked list has even number of elements, fast.Next will be nil
	// In both cases, slow will be pointing to the middle element
	// This is because fast will be pointing to the last element
	// and slow will be pointing to the middle element
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// If the linked list has even number of elements, fast.Next will be nil
	// In this case, slow will be pointing to the second middle element
	if fast.Next != nil {
		slow = slow.Next
	}
	return slow.Value, slow
}
