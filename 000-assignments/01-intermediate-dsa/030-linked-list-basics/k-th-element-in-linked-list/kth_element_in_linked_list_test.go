package kthelementinlinkedlist

import (
	ll "linkedlist"
	"testing"
)

func TestKthElementInLinkedList(t *testing.T) {
	type args struct {
		A *ll.Node
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{ll.NewNode(1, 2, 3), 0}, 1},
		{"test2", args{ll.NewNode(4, 3, 2, 1), 1}, 3},
		{"test3", args{ll.NewNode(11, 19, 11, 9, 1, 2), 5}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KthElementInLinkedList(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("KthElementInLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkKthElementInLinkedList-8   	12690448	       101.9 ns/op	      48 B/op	       3 allocs/op
func BenchmarkKthElementInLinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		KthElementInLinkedList(ll.NewNode(1, 2, 3), 0)
	}
}
