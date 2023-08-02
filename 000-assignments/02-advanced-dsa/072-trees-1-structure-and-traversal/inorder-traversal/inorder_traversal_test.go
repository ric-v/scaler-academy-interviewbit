package inordertraversal

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	type args struct {
		A *treeNode
	}

	tree1 := treeNode_new(1)
	tree1.right = treeNode_new(2)
	tree1.right.left = treeNode_new(3)

	tree2 := treeNode_new(1)
	tree2.left = treeNode_new(6)
	tree2.right = treeNode_new(2)
	tree2.right.left = treeNode_new(3)

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{tree1}, []int{1, 3, 2}},
		{"test2", args{tree2}, []int{6, 1, 3, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InorderTraversal(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkInorderTraversal/test1-8         	12759631	        95.41 ns/op	      56 B/op	       3 allocs/op
func BenchmarkInorderTraversal(b *testing.B) {
	tree1 := treeNode_new(1)
	tree1.right = treeNode_new(2)
	tree1.right.left = treeNode_new(3)

	tree2 := treeNode_new(1)
	tree2.left = treeNode_new(6)
	tree2.right = treeNode_new(2)
	tree2.right.left = treeNode_new(3)

	benchmarks := []struct {
		name string
		args *treeNode
	}{
		{"test1", tree1},
		{"test2", tree2},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				InorderTraversal(bm.args)
			}
		})
	}
}
