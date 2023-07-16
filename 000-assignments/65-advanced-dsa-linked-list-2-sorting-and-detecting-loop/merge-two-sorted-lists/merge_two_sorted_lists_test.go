package mergetwosortedlists

import (
	ll "linkedlist"
	"reflect"
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {
	type args struct {
		A *ll.Node
		B *ll.Node
	}
	tests := []struct {
		name string
		args args
		want *ll.Node
	}{
		{"test1", args{ll.NewNode(5, 8, 20), ll.NewNode(4, 11, 15)}, ll.NewNode(4, 5, 8, 11, 15, 20)},
		{"test2", args{ll.NewNode(1, 2, 3), nil}, ll.NewNode(1, 2, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeTwoSortedLists(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeTwoSortedLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMergeTwoSortedLists(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeTwoSortedLists(ll.NewNode(5, 8, 20), ll.NewNode(4, 11, 15))
	}
}
