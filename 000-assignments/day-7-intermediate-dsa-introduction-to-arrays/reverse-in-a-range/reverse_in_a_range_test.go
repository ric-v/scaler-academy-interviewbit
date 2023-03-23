package reverseinarange

import (
	"reflect"
	"testing"
)

func TestReverseInARange(t *testing.T) {
	type args struct {
		A []int
		B int
		C int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 2, 3, 4}, 2, 3}, []int{1, 2, 4, 3}},
		{"test2", args{[]int{2, 5, 6}, 0, 2}, []int{6, 5, 2}},
		{"test3", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, 8}, []int{9, 8, 7, 6, 5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseInARange(tt.args.A, tt.args.B, tt.args.C); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseInARange() = %v, want %v", got, tt.want)
			}
		})
	}
}
