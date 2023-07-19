package main

import (
	"fmt"
	ll "linkedlist"
)

func PrintLinkedList(A *ll.Node) {
	var temp *ll.Node = A
	for temp != nil {
		fmt.Print(temp.Value, " ")
		temp = temp.Next
	}
	fmt.Println()
}

func main() {
	PrintLinkedList(ll.NewNode(1, 2, 3, 4, 5))
	PrintLinkedList(ll.NewNode(4, 3, 2, 1))
}
