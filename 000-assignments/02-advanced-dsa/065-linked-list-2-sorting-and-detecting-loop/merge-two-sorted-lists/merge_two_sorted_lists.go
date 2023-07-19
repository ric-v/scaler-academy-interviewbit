package mergetwosortedlists

import (
	ll "linkedlist"
)

// MergeTwoSortedLists merges two sorted linked lists
// 
// Time complexity: O(n)
// Space complexity: O(1)
func MergeTwoSortedLists(A, B *ll.Node) *ll.Node {

	// If both linked lists are empty, return nil
	if A == nil {
		return B
	}
	if B == nil {
		return A
	}

	var result *ll.Node
	// If both linked lists are not empty, compare the first elements of both linked lists
	if A.Value < B.Value {
		result = A
		A = A.Next
	} else {
		result = B
		B = B.Next
	}

	// temp will be used to traverse the result linked list
	var temp *ll.Node = result

	// Traverse both linked lists and compare the elements
	// Append the smaller element to the result linked list
	// Update the temp pointer
	for A != nil && B != nil {
		if A.Value <= B.Value {
			temp.Next = A
			A = A.Next
		} else {
			temp.Next = B
			B = B.Next
		}
		temp = temp.Next
	}

	// If any of the linked lists is not empty, append it to the result linked list
	if A != nil {
		temp.Next = A
	}
	if B != nil {
		temp.Next = B
	}
	return result
}
