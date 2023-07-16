package searchinlinkedlist

import (
	ll "linkedlist"
	"testing"
)

func TestSearchInLinkedList(t *testing.T) {
	type args struct {
		A *ll.Node
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{ll.NewNode(1, 2, 3), 4}, 0},
		{"test2", args{ll.NewNode(4, 3, 2, 1), 4}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInLinkedList(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("SearchInLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSearchInLinkedList-8   	11382174	       101.1 ns/op	      48 B/op	       3 allocs/op
func BenchmarkSearchInLinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchInLinkedList(ll.NewNode(1, 2, 3), 4)
	}
}
