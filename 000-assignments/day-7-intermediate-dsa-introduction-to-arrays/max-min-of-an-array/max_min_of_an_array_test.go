package maxminofanarray

import "testing"

func TestMaxMinOfAnArray(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 3, 4, 5}}, 6},
		{"test2", args{[]int{5, 4, 3, 2, 1}}, 6},
		{"test3", args{[]int{1, 1, 1, 1, 1}}, 2},
		{"test4", args{[]int{8, 1, 2, 2, 3, 8, 4, 5, 7, 6}}, 9},
		{"test5", args{[]int{-2, 1, -4, 5, 3}}, 1},
		{"test6", args{[]int{1, 3, 4, 1}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxMinOfAnArray(tt.args.A); got != tt.want {
				t.Errorf("MaxMinOfAnArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
