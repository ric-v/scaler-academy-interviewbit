package balancedbinarytree

import "testing"

func TestBalancedBinaryTree(t *testing.T) {
	type args struct {
		A *treeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{treeNode_new(1)}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BalancedBinaryTree(tt.args.A); got != tt.want {
				t.Errorf("BalancedBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
