package levelorder

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	type args struct {
		A *treeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test1", args{treeNode_new(1)}, [][]int{{1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LevelOrder(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
