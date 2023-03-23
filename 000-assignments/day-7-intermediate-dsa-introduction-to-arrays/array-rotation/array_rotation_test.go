package arrayrotation

import (
	"reflect"
	"testing"
)

func TestArrayRotation(t *testing.T) {
	type args struct {
		A []int
		B int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// {"test1", args{[]int{1, 2, 3, 4, 5, 6}, 1}, []int{2, 3, 4, 5, 6, 1}},
		// {"test2", args{[]int{1, 2, 3, 4, 5, 6}, 2}, []int{3, 4, 5, 6, 1, 2}},
		// {"test3", args{[]int{1, 2, 3, 4, 5, 6}, 3}, []int{4, 5, 6, 1, 2, 3}},
		// {"test4", args{[]int{1, 2, 3, 4, 5, 6}, 4}, []int{5, 6, 1, 2, 3, 4}},
		// {"test5", args{[]int{1, 2, 3, 4, 5, 6}, 5}, []int{6, 1, 2, 3, 4, 5}},
		// {"test6", args{[]int{1, 2, 3, 4, 5, 6}, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"test7", args{[]int{1, 2, 3, 4}, 2}, []int{3, 4, 1, 2}},
		{"test8", args{[]int{2, 5, 6}, 1}, []int{6, 2, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayRotation(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}
