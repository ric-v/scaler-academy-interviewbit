package binarytreefrominorderandpostorder

import (
	"reflect"
	"testing"
)

func TestBinaryTreeFromInorderAndPostorder(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want *treeNode
	}{
		{"test1", args{[]int{2, 1, 3}, []int{2, 3, 1}}, &treeNode{left: &treeNode{value: 2}, value: 1, right: &treeNode{value: 3}}},
		{"test2", args{[]int{6, 1, 3, 2}, []int{6, 3, 2, 1}}, &treeNode{value: 1, left: &treeNode{value: 6}, right: &treeNode{value: 2, left: &treeNode{value: 3}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinaryTreeFromInorderAndPostorder(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryTreeFromInorderAndPostorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkBinaryTreeFromInorderAndPostorder-8   	 7196150	       165.3 ns/op	      96 B/op	       4 allocs/op
func BenchmarkBinaryTreeFromInorderAndPostorder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinaryTreeFromInorderAndPostorder([]int{6, 1, 3, 2}, []int{6, 3, 2, 1})
	}
}
