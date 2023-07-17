package removeloopfromlinkedlist

type ListNode struct {
	value int
	next  *ListNode
}

func ListNode_new(val int) *ListNode {
	var node *ListNode = new(ListNode)
	node.value = val
	node.next = nil
	return node
}

// RemoveLoopFromLinkedList removes loop from linked list
//
// Time complexity: O(n)
// Space complexity: O(1)
func RemoveLoopFromLinkedList(A *ListNode) *ListNode {
	var loopFound bool

	slow, fast := A, A
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
		if slow == fast {
			loopFound = true
			break
		}
	}
	if !loopFound {
		return A
	}
	slow = A
	// gives the start of cyle
	for slow != fast {
		slow = slow.next
		fast = fast.next
	}
	// by now slow and fast are at the start of cycle
	// let's find end of the cycle
	for fast.next != slow {
		fast = fast.next
	}
	fast.next = nil
	return A
}
