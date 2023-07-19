package insertinlinkedlist

import (
	"fmt"
	"linkedlist"
	"testing"
)

func TestInsertInLinkedList(t *testing.T) {
	type args struct {
		A *linkedlist.Node
		B int
		C int
	}
	tests := []struct {
		name string
		args args
		want *linkedlist.Node
	}{
		{"test1", args{linkedlist.NewNode(1, 2), 3, 0}, linkedlist.NewNode(3, 1, 2)},
		{"test2", args{linkedlist.NewNode(1, 2), 3, 1}, linkedlist.NewNode(1, 3, 2)},
		{"test3", args{linkedlist.NewNode(1, 2, 3, 4, 5), 6, 3}, linkedlist.NewNode(1, 2, 3, 6, 4, 5)},
		{"test4", args{linkedlist.NewNode(1, 2), 3, 2}, linkedlist.NewNode(1, 2, 3)},
		{"test5", args{linkedlist.NewNode(7), 5, 7}, linkedlist.NewNode(7, 5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertInLinkedList(tt.args.A, tt.args.B, tt.args.C); !got.Equals(tt.want) {
				fmt.Println(got, tt.want)
				t.Errorf("InsertInLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkInsertInLinkedList-8   	11620789	       103.2 ns/op	      48 B/op	       3 allocs/op
func BenchmarkInsertInLinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertInLinkedList(linkedlist.NewNode(1, 2), 3, 0)
	}
}
