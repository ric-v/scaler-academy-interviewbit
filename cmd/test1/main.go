package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func listNode_new(val int) *ListNode {
	var node *ListNode = new(ListNode)
	node.val = val
	node.next = nil
	return node
}

func main() {

	var A *ListNode = listNode_new(1)
	A.next = listNode_new(4)
	A.next.next = listNode_new(5)
	A.next.next.next = listNode_new(6)
	A.next.next.next.next = listNode_new(7)
	res := solve(A)
	for res != nil {
		fmt.Print(res.val, " -> ")
		res = res.next
	}

}

// remove linkedlist elements that are odd numbers
func solve(A *ListNode) *ListNode {

	var head *ListNode = A
	var prev *ListNode = nil
	var curr *ListNode = A

	for curr != nil {
		if curr.val%2 != 0 {
			if prev == nil {
				head = curr.next
			} else {
				prev.next = curr.next
			}
		} else {
			prev = curr
		}
		curr = curr.next
	}

	return head
}
