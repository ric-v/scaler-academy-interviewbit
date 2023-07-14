package palindromelist

import (
	ll "linkedlist"
)

func PalindromeList(A *ll.Node) int {
	p1 := A
	p2 := A
	for {
		if p2.Next != nil && p2.Next.Next == nil {
			break
		} else if p2.Next == nil {
			break
		}
		p2 = p2.Next.Next
		p1 = p1.Next
	}
	mid := p1.Next
	prev := (*ll.Node)(nil)
	for mid != nil {
		c := mid.Next
		mid.Next = prev
		prev = mid
		mid = c
	}
	mid = prev
	temp := A
	for mid != nil && temp != nil {
		if temp.Value != mid.Value {
			return 0
		}
		mid = mid.Next
		temp = temp.Next
	}
	return 1
}
