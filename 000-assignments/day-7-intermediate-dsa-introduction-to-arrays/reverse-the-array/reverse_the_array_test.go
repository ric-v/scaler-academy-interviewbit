package reversethearray

import (
	"reflect"
	"testing"
)

func TestReverseTheArray(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 2, 3, 4}}, []int{4, 3, 2, 1}},
		{"test2", args{[]int{2, 5, 6}}, []int{6, 5, 2}},
		{"test3", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}}, []int{9, 8, 7, 6, 5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseTheArray(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseTheArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
