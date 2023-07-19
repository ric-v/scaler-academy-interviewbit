package middleelementoflinkedlist

import (
	ll "linkedlist"
	"testing"
)

func TestMiddleElementOfLinkedList(t *testing.T) {
	type args struct {
		A *ll.Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{ll.NewNode(1, 2, 3, 4, 5)}, 3},
		{"test2", args{ll.NewNode(1, 5, 6, 2, 3, 4)}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := MiddleElementOfLinkedList(tt.args.A); got != tt.want {
				t.Errorf("MiddleElementOfLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkMiddleElementOfLinkedList-8   	 7120414	       171.6 ns/op	      80 B/op	       5 allocs/op
func BenchmarkMiddleElementOfLinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MiddleElementOfLinkedList(ll.NewNode(1, 2, 3, 4, 5))
	}
}
