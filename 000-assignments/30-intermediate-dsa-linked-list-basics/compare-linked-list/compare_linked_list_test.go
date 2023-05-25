package comparelinkedlist

import (
	ll "linkedlist"
	"testing"
)

func TestCompareLinkedList(t *testing.T) {
	type args struct {
		A *ll.Node
		B *ll.Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{ll.NewNode(1, 2, 3, 4), ll.NewNode(1, 2, 3, 4)}, 1},
		{"test2", args{ll.NewNode(1, 2, 3, 4), ll.NewNode(1, 2, 3, 4, 5)}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareLinkedList(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("CompareLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkCompareLinkedList-8   	 3445980	       408.9 ns/op	     128 B/op	       8 allocs/op
func BenchmarkCompareLinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompareLinkedList(ll.NewNode(1, 2, 3, 4), ll.NewNode(1, 2, 3, 4))
	}
}
