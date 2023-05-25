package linkedlist

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func NewNode(val ...int) *Node {
	var node *Node = new(Node)
	node.Value = val[0]
	if len(val) > 1 {
		node.Next = NewNode(val[1:]...)
	}
	return node
}

func (n *Node) Equals(m *Node) bool {
	if n == nil && m == nil {
		return true
	}
	if n == nil || m == nil {
		return false
	}
	return n.Value == m.Value && n.Next.Equals(m.Next)
}

func (n *Node) String() string {
	if n == nil {
		return ""
	}
	if n.Next == nil {
		return fmt.Sprint(n.Value)
	}
	return fmt.Sprint(n.Value) + " " + n.Next.String()
}
