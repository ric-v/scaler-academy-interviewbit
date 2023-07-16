package sortlinkedlist

import (
	ll "linkedlist"
	"reflect"
	"testing"
)

func TestSortLinkedList(t *testing.T) {
	type args struct {
		A *ll.Node
	}
	tests := []struct {
		name string
		args args
		want *ll.Node
	}{
		{"test1", args{ll.NewNode(3, 4, 2, 8)}, ll.NewNode(2, 3, 4, 8)},
		{"test2", args{ll.NewNode(1)}, ll.NewNode(1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortLinkedList(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSortLinkedList-8   	 7520817	       158.4 ns/op	      64 B/op	       4 allocs/op
func BenchmarkSortLinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SortLinkedList(ll.NewNode(3, 4, 2, 8))
	}
}
