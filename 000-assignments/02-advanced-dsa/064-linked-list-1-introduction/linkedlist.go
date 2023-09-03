package main

import (
	"fmt"
	"linkedlist"
)

func main() {
	ll := linkedlist.NewNode(1, 2, 3, 4, 5, 6)
	fmt.Println(ll.String())

	fmt.Println(search(ll, 3))
	fmt.Println(search(ll, 7))

	ll = insert(ll, 0, 10)
	fmt.Println(ll.String())

	ll = insert(ll, 3, 20)
	fmt.Println(ll.String())

	ll = delete(ll, 10)
	fmt.Println(ll.String())

	ll = delete(ll, 20)
	fmt.Println(ll.String())

	ll = reverse(ll)
	fmt.Println(ll.String())
}

func search(ll *linkedlist.Node, k int) bool {

	if ll == nil {
		return false
	}

	temp := ll

	for temp.Next != nil {

		if temp.Value == k {
			return true
		}
		temp = temp.Next
	}
	return false
}

func insert(ll *linkedlist.Node, k int, v int) *linkedlist.Node {

	newLL := linkedlist.NewNode(v)

	if k == 0 {
		newLL.Next = ll
		ll = newLL
	} else {

		temp := ll
		for i := 0; i < k-1; i++ {
			temp = temp.Next
		}

		newLL.Next = temp.Next
		temp.Next = newLL
	}
	return ll
}

func delete(ll *linkedlist.Node, k int) *linkedlist.Node {

	if ll == nil {
		return ll
	}

	temp := ll
	if temp.Value == k {
		ll = temp.Next
		return ll
	}

	for temp != nil {
		if temp.Next != nil && temp.Next.Value == k {
			temp.Next = temp.Next.Next
			return ll
		}
		temp = temp.Next
	}
	return ll
}

func reverse(ll *linkedlist.Node) *linkedlist.Node {

	var prev, curr *linkedlist.Node = nil, ll

	fmt.Println(min(1, 2))

	for curr != nil {

		forward := curr.Next
		curr.Next = prev
		prev = curr
		curr = forward
	}
	ll = prev
	return ll
}
