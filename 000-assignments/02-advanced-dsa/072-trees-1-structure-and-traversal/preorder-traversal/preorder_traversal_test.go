package preordertraversal

import (
	"reflect"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
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
		{"test1", args{tree1}, []int{1, 2, 3}},
		{"test2", args{tree2}, []int{1, 6, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PreorderTraversal(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkPreorderTraversal/test1-8         	11657773	        95.15 ns/op	      56 B/op	       3 allocs/op
func BenchmarkPreorderTraversal(b *testing.B) {
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
				PreorderTraversal(bm.args)
			}
		})
	}
}
