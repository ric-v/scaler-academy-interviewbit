package reverselinkedlist

import (
	ll "linkedlist"
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	type args struct {
		A *ll.Node
	}
	tests := []struct {
		name string
		args args
		want *ll.Node
	}{
		{"test1", args{ll.NewNode(1, 2, 3, 4, 5)}, ll.NewNode(5, 4, 3, 2, 1)},
		{"test2", args{ll.NewNode(5, 4, 3, 2, 1)}, ll.NewNode(1, 2, 3, 4, 5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseList(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkReverseList-8   	 7093837	       170.9 ns/op	      80 B/op	       5 allocs/op
func BenchmarkReverseList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseList(ll.NewNode(1, 2, 3, 4, 5))
	}
}
