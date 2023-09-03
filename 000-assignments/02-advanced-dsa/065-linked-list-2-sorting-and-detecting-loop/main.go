package main

import (
	"fmt"
	"strings"
)

type Node struct {
	data int
	next *Node
}

func NewNode(data int) *Node {
	return &Node{data: data}
}

func Append(head *Node, data int) {
	node := NewNode(data)
	if head == nil {
		head = node
		return
	}
	for head.next != nil {
		head = head.next
	}
	head.next = node
}

func (n *Node) String() string {
	var s strings.Builder
	for n != nil {
		s.WriteString(fmt.Sprintf("%v -> ", n.data))
		n = n.next
	}
	s.WriteString("nil")
	return s.String()
}

func (n *Node) Mid() *Node {
	if n == nil || n.next == nil {
		return n
	}
	slow, fast := n, n
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

func (n *Node) merge(ll2 *Node) *Node {
	if n == nil {
		return ll2
	}
	if ll2 == nil {
		return n
	}

	var head *Node

	if n.data <= ll2.data {
		head = n
		n = n.next
	} else {
		head = ll2
		ll2 = ll2.next
	}

	var curr = head
	for n != nil && ll2 != nil {
		if n.data < ll2.data {
			curr.next = n
			n = n.next
		} else {
			curr.next = ll2
			ll2 = ll2.next
		}
		curr = curr.next
	}

	if n != nil {
		curr.next = n
	}
	if ll2 != nil {
		curr.next = ll2
	}

	return head
}

func (n *Node) mergeSort() *Node {

	if n == nil || n.next == nil {
		return n
	}

	mid := n.Mid()
	h2 := mid.next
	mid.next = nil

	left := n.mergeSort()
	right := h2.mergeSort()

	return left.merge(right)
}

func main() {

	head := NewNode(5)
	Append(head, 1)
	Append(head, 3)
	Append(head, 2)
	Append(head, 4)

	fmt.Println(head)

	mid := head.Mid()
	fmt.Println(mid.data)

	head1 := NewNode(1)
	Append(head1, 2)
	Append(head1, 3)
	Append(head1, 4)
	Append(head1, 5)
	fmt.Println(head1)

	head2 := NewNode(6)
	Append(head2, 7)
	Append(head2, 8)
	Append(head2, 9)
	Append(head2, 10)
	fmt.Println(head2)

	head = head1.merge(head2)
	fmt.Println(head)

	head = NewNode(5)
	Append(head, 1)
	Append(head, 3)
	Append(head, 2)
	Append(head, 4)
	fmt.Println(head)

	head = head.mergeSort()
	fmt.Println(head)

}
