package sortlinkedlist

import (
	ll "linkedlist"
	mergetwosortedlists "merge-two-sorted-listsmerge-two-sorted-lists"
)

func SortLinkedList(A *ll.Node) *ll.Node {

	if A == nil || A.Next == nil {
		return A
	}

	// Find the middle of the linked list
	var mid = findMid(A)
	var midNext = mid.Next
	mid.Next = nil

	// Sort the first half of the linked list
	var left = SortLinkedList(A)

	// Sort the second half of the linked list
	var right = SortLinkedList(midNext)

	// Merge the sorted linked lists
	return mergetwosortedlists.MergeTwoSortedLists(left, right)
}

func findMid(A *ll.Node) *ll.Node {
	// If the linked list is empty, return 0
	if A == nil {
		return nil
	}

	var slow, fast *ll.Node = A, A
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
